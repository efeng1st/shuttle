package log

import (
	"errors"
	"time"
)

const (
	LogModeOff     = "off"
	LogModeConsole = "console"
	LogModeFile    = "file"
)

type ILogConfig interface {
	GetLogLevel() string
}

func InitLogger(logMode, logPath string) (err error) {
	switch logMode {
	case LogModeOff:
		Logger, err = NewSkipLogger()
		if err != nil {
			return errors.New("init logger failed:" + err.Error())
		}
	case LogModeConsole:
		Logger, err = NewStdLogger(LogInfo)
		if err != nil {
			return errors.New("init logger failed:" + err.Error())
		}
	case LogModeFile:
		//multiSize: 100MB
		Logger, err = NewFileLogger(logPath, LogInfo, 100*1000*1000)
		if err != nil {
			return errors.New("init logger failed:" + err.Error())
		}
	default:
		return errors.New("not support LogMode:" + logMode)
	}
	return
}

func ApplyConfig(logConfig ILogConfig) error {
	levelFlag, ok := LevelMap[logConfig.GetLogLevel()]
	if !ok {
		return errors.New("not support LogLevel:" + logConfig.GetLogLevel())
	}
	Logger.SetLevel(levelFlag)
	return nil
}

var Logger ILogger = &StdLogger{Level: LogDebug}

func SetLogger(logger ILogger) {
	Logger = logger
}

const (
	LogTrace = 0
	LogDebug = 1
	LogInfo  = 2
	LogError = 3
)

var LevelMap = map[string]int{
	"trace": LogTrace,
	"debug": LogDebug,
	"info":  LogInfo,
	"error": LogError,
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

type ILogger interface {
	SetLevel(int)
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Error(...interface{})
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Close() error
}
