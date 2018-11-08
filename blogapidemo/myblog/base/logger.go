package base

import (
	"os"
	"fmt"
	"myblog/conf"
	"github.com/op/go-logging"
	"io"
)

var Logger * logging.Logger
var LoggerWriter io.Writer

func InitLogger()error{
	configData  := conf.Config
	ok, _ := PathExists(configData.LogDir)
	if !ok {
		err := os.MkdirAll(configData.LogDir, 0777)
		if nil != err {
			fmt.Println("can't make dir : %s, %v", conf.Config.LogDir, err)
			return err
		}
	}

	LogFormat := "%{color}%{level:.4s}:%{time:2006-01-02 15:04:05.000}[%{id:03x}] %{shortfile} %{color:reset} %{message}"
	Logger = logging.MustGetLogger("")

	writeFd, err := os.OpenFile(configData.LogDir+"/"+configData.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file[%s] failed[%s]", configData.LogFile, err)
		return err
	}
	LoggerWriter = writeFd


	format := logging.MustStringFormatter(LogFormat)
	backendList := make([] logging.Backend, 0)
	for _, out := range [] io.Writer{writeFd, os.Stdout}{
		backendInfo := logging.NewLogBackend(out, "", 0)
		backendInfoFormatter := logging.NewBackendFormatter(backendInfo, format)
		backendInfoLeveld := logging.AddModuleLevel(backendInfoFormatter)

		switch configData.LogLevel {
		case "ERROR":
			backendInfoLeveld.SetLevel(logging.ERROR, "")
		case "WARNING":
			backendInfoLeveld.SetLevel(logging.WARNING, "")
		case "NOTICE":
			backendInfoLeveld.SetLevel(logging.NOTICE, "")
		case "INFO":
			backendInfoLeveld.SetLevel(logging.INFO, "")
		case "DEBUG":
			backendInfoLeveld.SetLevel(logging.DEBUG, "")
		default:
			backendInfoLeveld.SetLevel(logging.ERROR, "")
		}
		bk := backendInfoLeveld
		backendList = append(backendList, bk)
	}

	logging.SetBackend(backendList...)

	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}