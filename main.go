package main

import (
	"fmt"
	"github.com/ljzh152535/framework/goadmin-consul"
)

func main() {
	//
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.10|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.10:9121"
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.12|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.12:9121"
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.13|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.13:9122"
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.8|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.8:9122"
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.9|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.9:9121"
	//sh 1.sh -e "${tmp_consul_ip}|redis|10.116.7.11|2000|cp-redis-idgen|zhuzhoujinke|10.116.7.11:9122"

	register, err := goadmin_consul.NewConuApi("redis", "172.30.2.106", 6679,
		"FireflyRedis", "uat", "172.30.2.106:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.113", 6679,
		"FireflyRedis", "uat", "172.30.2.113:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.114", 6679,
		"FireflyRedis", "uat", "172.30.2.114:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.115", 6680,
		"FireflyRedis", "uat", "172.30.2.115:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.116", 6680,
		"FireflyRedis", "uat", "172.30.2.116:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.117", 6680,
		"FireflyRedis", "uat", "172.30.2.117:9121", "uat").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.106", 6679,
		"FireflyAPPREDIS", "uat", "172.30.2.106:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.113", 6679,
		"FireflyAPPREDIS", "uat", "172.30.2.113:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.114", 6679,
		"FireflyAPPREDIS", "uat", "172.30.2.114:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.115", 6680,
		"FireflyAPPREDIS", "uat", "172.30.2.115:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.116", 6680,
		"FireflyAPPREDIS", "uat", "172.30.2.116:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}

	register, err = goadmin_consul.NewConuApi("redis", "172.30.2.117", 6680,
		"FireflyAPPREDIS", "uat", "172.30.2.117:9121", "hexi").Register()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(register)
	}
}
