package logger

import (
	"fmt"
	"log"
	"loganalysis/config"
	"os"
	"path/filepath"
	"time"
)

func init() {
	logRootPath := config.Config.LogPath
	logPath := filepath.Join(logRootPath, time.Now().Format("2006"), "/", time.Now().Format("01"), time.Now().Format("02"), "/")
	logName := "log.txt"
	//检查文件夹是否存在 不存在则创建
	_, err := os.Stat(logPath)
	if err != nil {
		err = os.MkdirAll(logPath, 0777)
		if err != nil {
			fmt.Println("日志文件夹创建失败，程序退出", err.Error())
			os.Exit(1)
		}
	}
	logFile, err := os.OpenFile(filepath.Join(logPath, logName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}
