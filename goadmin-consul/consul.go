package goadmin_consul

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	//consulApi "github.com/armon/consul-api"
	"strconv"
)

type ConulRegisterConfig struct {
	ID                string
	Name              string
	Address           string
	Port              int
	Tags              []string
	Meta              map[string]string
	HTTP              string
	Interval          string
	redisExporterAddr string
}

func NewConuApi(consulType string, redisIp string, redisPort int, appName string, zone string, redisExporterAddr string, cc string) *ConulRegisterConfig {
	return &ConulRegisterConfig{
		ID:      consulType + "_" + appName + "_" + redisIp + "_" + strconv.Itoa(redisPort),
		Name:    "consul_" + consulType + "_" + appName,
		Tags:    []string{redisExporterAddr, consulType},
		Address: redisIp,
		Port:    redisPort,
		Meta:    map[string]string{"appName": appName, "zone": zone}, //"consul.datacenter": "hexi",

		redisExporterAddr: fmt.Sprintf("http://%s%s", redisExporterAddr, "/metrics"),
	}
}

func (c *ConulRegisterConfig) Register() (string, error) {
	config := consulApi.DefaultConfig()
	config.Address = "172.30.2.99:15500"
	config.Datacenter = "zhezhong"

	consulClient, err := consulApi.NewClient(config)
	if err != nil {
		return "", fmt.Errorf("register server failed: [%w]", err)
	}

	//reg := &consulApi.CatalogRegistration{
	//	Address: "172.30.1.221",
	//	Node:    "consul-01",
	//	//Datacenter: "zhezhong",
	//	Service: &consulApi.AgentService{
	//		ID:                c.ID,
	//		Service:           c.Name,
	//		Address:           c.Address,
	//		Port:              c.Port,
	//		Tags:              c.Tags,
	//		Meta:              c.Meta,
	//		EnableTagOverride: false,
	//	},
	//	//NodeMeta: map[string]string{"dc": "zhezhong"},
	//}

	//register, err := consulClient.Catalog().Register(reg, nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(register)

	// 注册信息
	registration := new(consulApi.AgentServiceRegistration)
	registration.ID = c.ID
	registration.Name = c.Name
	registration.Address = c.Address
	registration.Port = c.Port
	registration.Tags = c.Tags
	registration.Meta = c.Meta
	registration.EnableTagOverride = false

	//增加check。
	check := new(consulApi.AgentServiceCheck)
	check.HTTP = c.redisExporterAddr
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "10s"
	//注册check服务。
	registration.Check = check
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		return "", fmt.Errorf("register server failed: [%w]", err)
	}
	return fmt.Sprintf("%s:%s", c.ID, "注册成功"), nil
}

//// 注册服务
//// consulClient.Agent()先获取当前机器上的consul agent节点
//consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
//	ID:      "MyService",
//	Name:    "My Service",
//	Address: "127.0.0.1",
//	Port:    5050,
//	Check: &api.AgentServiceCheck{
//		CheckID:  "MyService",
//		TCP:      "127.0.0.1:5050",
//		Interval: "10s",
//		Timeout:  "1s",
//	},
//})
//
//// 运行完成后注销服务
//defer consulClient.Agent().ServiceDeregister("MyService")
//l, err := net.Listen("tcp", ":5050")
//if err != nil {
//	log.Fatal(err)
//}
//for {
//	conn, err := l.Accept()
//	if err != nil {
//		log.Fatal(err)
//	}
//	go func() {
//		log.Printf("Ip: %s connected", conn.RemoteAddr().String())
//	}()
//}
