package config

import (
	goadmin_config "github.com/ljzh152535/framework/goadmin-config"
	goadmin_dbmodel "github.com/ljzh152535/framework/goadmin-db/model"
	goadmin_logrus "github.com/ljzh152535/framework/goadmin-logrus"
)

type Config struct {
	LOG          goadmin_logrus.CoreLogrus         `mapstructure:"log" yaml:"log"`
	DB           map[string]goadmin_dbmodel.DBItem `yaml:"db"`
	System       goadmin_config.System             `mapstructure:"system" json:"system" yaml:"system"`
	WebServerLog goadmin_config.WebServerLog       `yaml:"web_server_log" mapstructure:"web_server_log"`
}
