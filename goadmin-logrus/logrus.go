package goadmin_logrus

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// log实例
const logTimeTpl = "2006-01-02T15:04:05.000Z07:00"

var log *logrus.Entry
var logLocker sync.RWMutex

type CoreLogrus struct {
	Level             string `mapstructure:"level" json:"level" yaml:"level"`                                     // 级别
	LogFormat         string `mapstructure:"logFormat" json:"logFormat" yaml:"logFormat"`                         // 日志格式 json text
	IsSetReportCaller bool   `mapstructure:"isSetReportCaller" json:"isSetReportCaller" yaml:"isSetReportCaller"` // 显示文件和代码行数
	LogEnv            string `mapstructure:"logEnv" json:"logEnv" yaml:"logEnv"`                                  // 日志系统环境
	Output            string `mapstructure:"output" json:"output" yaml:"output"`                                  // 日志输出方式  file output
	LogPath           string `mapstructure:"logPath" json:"logPath" yaml:"logPath"`                               // 日志路径
	MaxAge            int64  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                               // 日志留存时间
	MaxCapacity       int64  `mapstructure:"max-capacity" json:"max-capacity" yaml:"max-capacity"`                // 单个日志的最大容量
	HostName          string `mapstructure:"hostName" json:"hostName" yaml:"hostName"`                            // 主机名
}

func InitLogrus(logConf CoreLogrus) *logrus.Entry {
	logLocker.RLock()
	if log != nil {
		logLocker.RUnlock()
		return log
	}
	logLocker.RUnlock()

	// A,B,C
	logLocker.Lock()
	defer logLocker.Unlock()

	// 二次判断
	if log != nil {
		return log
	}

	logNew := logrus.New()

	// 设置log的配置
	return setLogrusConf(logConf, logNew)
}

// 设置日志level
func setLogrusLevel(logConf CoreLogrus, logNew *logrus.Logger) {
	switch logConf.Level {
	case "debug":
		logNew.SetLevel(logrus.DebugLevel)
	case "error":
		logNew.SetLevel(logrus.ErrorLevel)
	case "warn":
		logNew.SetLevel(logrus.WarnLevel)
	default:
		logNew.SetLevel(logrus.InfoLevel)
	}
}

// 设置日志格式
func setLogrusFormat(logConf CoreLogrus, logNew *logrus.Logger, entry *logrus.Entry) {
	if logConf.LogFormat == "json" {
		logNew.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: logTimeTpl,
			// runtime.Frame: 帧,可用于获取调用者返回的PC值的函数、文件或者是行信息
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fileName := path.Base(frame.File)
				return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
			},
		})
		logNew.SetOutput(setRotatelogs(logConf))
	} else {
		// 如果非dev环境禁用掉color
		logNew.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: logTimeTpl,
			DisableColors:   logConf.LogEnv != "dev",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fName := filepath.Base(entry.Caller.File)
				return frame.Function, fmt.Sprintf("[%s] [%-7s] [%s:%d %s] %s\n",
					logTimeTpl, entry.Level.String(), fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
			},
		})
	}
}

// 设置日志输出方式
func setLogrusOutput(logConf CoreLogrus, logNew *logrus.Logger) {
	if logConf.Output == "file" {
		f, e := loadLogFile(logConf)
		if e != nil {
			panic(e)
		}
		logNew.SetOutput(f)
	} else {
		logNew.SetOutput(os.Stdout)
	}
}

func loadLogFile(logConf CoreLogrus) (io.Writer, error) {
	logPath := "logs/app.log"
	if logConf.LogPath != "" {
		logPath = logConf.LogPath
	}

	// 判断logPath是相对路径还是绝对路径
	//if !filepath.IsAbs(logPath) {
	//	logPath = homeDir + "/" + logPath
	//}

	// 检查文件是否存在，不存在创建文件
	f, e := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		return nil, e
	}
	return f, nil
}

func setLogrusConf(logConf CoreLogrus, logNew *logrus.Logger) *logrus.Entry {

	// 设置日志level
	setLogrusLevel(logConf, logNew)

	// 显示文件和代码行数
	logNew.SetReportCaller(logConf.IsSetReportCaller) // 显示文件和代码行数

	// 设置日志输出方式
	setLogrusOutput(logConf, logNew)

	// 基础字段预设,比如项目名、环境、env、local_ip、hostname、idc
	l := logNew.WithFields(logrus.Fields{
		"env": logConf.LogEnv,
		//"loccal_ip": env.LocalIP(),
		"hostname": logConf.HostName,
	})

	// 设置日志格式 json 或者 text
	setLogrusFormat(logConf, logNew, l)
	log = l
	return l
}

// 设置日志切割
func setRotatelogs(logConf CoreLogrus) *rotatelogs.RotateLogs {
	logfile := logConf.LogPath

	writer, err := rotatelogs.New(
		logfile+"-%Y%m%d.log",
		rotatelogs.WithLinkName(logfile), //生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(logConf.MaxAge)), // 最大的保留时间 单位: 天
		rotatelogs.WithRotationTime(24*time.Hour),                         //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		rotatelogs.WithRotationSize(logConf.MaxCapacity*1024*1024),        // 设置分割文件的大小为  单位: MB
	)

	if err != nil {
		log.Error(err)
	}
	return writer
}
