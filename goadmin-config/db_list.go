package goadmin_config

type GeneralDB struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Config       string `yaml:"config"`                                       // 高级配置
	Timeout      int    `yaml:"timeout"`                                      // connect db timeout , uint ms
	WriteTimeOut int    `yaml:"write_time_out" mapstructure:"write_time_out"` // write data timeout , uint ms
	ReadTimeOut  int    `yaml:"read_time_out" mapstructure:"read_time_out"`   // read data timeout,uint ms
	MaxIdleConns int    `yaml:"max_idle_conns" mapstructure:"max_idle_conns"` // 最大的闲置连接数
	MaxOpenConns int    `yaml:"max_open_conns" mapstructure:"max_open_conns"` //最大打开连接数
}

type DBItem struct {
	Write GeneralDB `yaml:"write"`
	Read  GeneralDB `yaml:"read"`
}
