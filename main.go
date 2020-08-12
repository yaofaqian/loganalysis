package main

import (
	"log"
	"loganalysis/config"
	"loganalysis/logic"
	"time"
)

func main() {
	startTime := time.Now()
	logTime := time.Now().Format("2006-01-02")
	if config.Config.LogPkgTime != "" {
		logTime = config.Config.LogPkgTime
	}
	LogAnalysisLogic := new(logic.LogAnalysisLogic)
	LogAnalysisLogic.SetDirPath(config.Config.DirPath)
	LogAnalysisLogic.SetLogTime(logTime)
	LogAnalysisLogic.SetHowManyInTheTop(config.Config.HowManyInTheTop)
	LogAnalysisLogic.SetGeneralLogChan(3000)
	LogAnalysisLogic.SetCloseProgramChan(10)
	LogAnalysisLogic.ZipDecompress()
	//util.LastAcquisitionDateAdd(logTime)
	//计算程序运行时间
	endTime := time.Since(startTime)
	log.Printf("程序执行完成,程序运行时间---%s", endTime)
}
