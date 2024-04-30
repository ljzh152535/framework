package model

type DBItemConf struct {
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

type DBLog struct {
	Enable bool   `yaml:"enable"`
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Type   string `yaml:"type"`
	Path   string `yaml:"path"`
}

type DBItem struct {
	Write DBItemConf `yaml:"write"`
	Read  DBItemConf `yaml:"read"`
	Log   DBLog      `yaml:"log"`
}
