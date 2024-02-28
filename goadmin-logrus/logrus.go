package goadmin_logrus

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// log实例

var (
	log        = logrus.New()
	coreLogrus CoreLogrus
)

// 初始化配置
func init() {
}

type logType string

const (
	A logType = "SetOutput"                       // 控制台 输出
	B logType = "SetOutputMultiWriterLogFile"     // 控制台 + 文件 输出
	C logType = "SetOutputJSON"                   // 控制台 json输出
	D logType = "SetOutputJSONMultiWriterLogFile" // 控制台 + 日志文件 json输出
	E logType = "SetWriterLogFile"                // 日志文件 输出
	F logType = "SetWriterJSONLogFile"            // 日志文件 json输出
)

type LogFormatter struct{}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var msg string
	//entry.Logger.SetReportCaller(true)
	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		msg = fmt.Sprintf("[%s] [%-7s] [%s:%d %s] %s\n",
			timestamp, entry.Level.String(), fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(msg)
	return b.Bytes(), nil
}

// 控制台输出
func setOutputCommon() *logrus.Logger {
	log.SetReportCaller(true) // 显示文件和代码行数
	log.SetOutput(os.Stdout)
	return log
}

// 控制台 + 日志文件 输出
func setOutputMultiWriterCommon(filePath string) *logrus.Logger {
	//f, _ := os.OpenFile(logConifg.filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log.SetReportCaller(true) // 显示文件和代码行数
	//log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetOutput(io.MultiWriter(setRotatelogs(filePath), os.Stdout))
	return log
}

// 日志文件 输出
func setWriterLogFileCommon(filePath string) *logrus.Logger {
	log.SetReportCaller(true) // 显示文件和代码行数
	log.SetOutput(setRotatelogs(filePath))
	return log
}

func setRotatelogs(filePath string) *rotatelogs.RotateLogs {
	logfile := filePath

	writer, err := rotatelogs.New(
		logfile+"-%Y%m%d.log",
		rotatelogs.WithLinkName(logfile+".log"),                              //生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(coreLogrus.MaxAge)), // 最大的保留时间 单位: 天
		rotatelogs.WithRotationTime(24*time.Hour),                            //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		rotatelogs.WithRotationSize(coreLogrus.MaxCapacity*1024*1024),        // 设置分割文件的大小为  单位: MB
	)

	if err != nil {
		log.Error(err)
	}
	return writer
}

func connditionSelect(typeName logType, logTypeName logType) {
	if typeName == logTypeName {
		log.SetFormatter(&LogFormatter{})
	} else {
		//log.SetFormatter(&logrus.JSONFormatter{})
		// 日志格式改成json
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			// runtime.Frame: 帧,可用于获取调用者返回的PC值的函数、文件或者是行信息
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fileName := path.Base(frame.File)
				return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
			},
		})
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
func transportLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.WarnLevel
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.DebugLevel
	}
}

type CoreLogrus struct {
	Level          string `mapstructure:"level" json:"level" yaml:"level"`                             // 级别
	Prefix         string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                          // 日志前缀
	Director       string `mapstructure:"director" json:"director"  yaml:"director"`                   // 日志文件夹
	PrefixFileName string `mapstructure:"prefixFileName" json:"prefixFileName"  yaml:"prefixFileName"` // 日志名前缀
	MaxAge         int64  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                       // 日志留存时间
	MaxCapacity    int64  `mapstructure:"max-capacity" json:"max-capacity" yaml:"max-capacity"`        // 单个日志的最大容量
}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func InitLogurs(a interface{}, logrusType logType, isGin bool) *logrus.Logger {
	v := reflect.ValueOf(a)
	coreLogrus.Level = v.FieldByName("Level").String()
	coreLogrus.Director = v.FieldByName("Director").String()
	coreLogrus.PrefixFileName = v.FieldByName("PrefixFileName").String()
	coreLogrus.MaxAge = v.FieldByName("MaxAge").Int()
	coreLogrus.MaxCapacity = v.FieldByName("MaxCapacity").Int()

	dir := coreLogrus.Director

	if ok, _ := pathExist(dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}
	//log := logrus.New()
	return newLogger(logrusType, coreLogrus, isGin)
}

// logrus *logrus.Logger
// typeName logType 打印类型
// a interface{} config.yaml logrus的配置
func newLogger(typeName logType, coreLogrus CoreLogrus, isGin bool) *logrus.Logger {

	//log = logrus
	var filePath string
	filePath = "logs/go-admin" // 默认值

	level := transportLevel(coreLogrus.Level)

	log.SetLevel(level)
	if isGin {
		ginMode := gin.Mode()
		if ginMode == gin.ReleaseMode {
			gin.SetMode(gin.ReleaseMode) // 线上模式，控制台不会打印信息
		}
		gin.DefaultWriter = log.Out // gin框架自己记录的日志也会输出
	}
	switch typeName {
	case A, C:
		connditionSelect(typeName, A)
		return setOutputCommon()
	case B, D:
		connditionSelect(typeName, B)
		filePath = coreLogrus.Director + "/" + coreLogrus.PrefixFileName
		return setOutputMultiWriterCommon(filePath)
	case E, F:
		connditionSelect(typeName, E)
		filePath = coreLogrus.Director + "/" + coreLogrus.PrefixFileName
		return setWriterLogFileCommon(filePath)
	default:
		log.Fatalf("输入的类型不对")
	}
	return nil
}
