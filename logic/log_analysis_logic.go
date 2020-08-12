package logic

import (
	"fmt"
	_ "github.com/lingdor/glog2midlog"
	"github.com/lingdor/midlog"
	"io/ioutil"
	"log"
	"loganalysis/antpool"
	_ "loganalysis/logger"
	"loganalysis/services"
	"loganalysis/util"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	GeneralLogAnalysis  = 1
	SecurityLogAnalysis = 2
	Collected           = 1
	NotCollected        = 0
)

//日志分析处理
type LogAnalysisLogic struct {
	dirPath         string //设置根路径
	logTime         string //设置哪天
	howManyInTheTop int    //所有的排行榜取前多少名
	generalLogChan  chan []string
	closeProgram    chan string
	taskCount       int //总任务数
}

var Logger = midlog.New("logic")
var logFileTypeList = map[int]string{1: "常规分析报告", 2: "安全分析报告"}

//设置文件夹路径
func (this *LogAnalysisLogic) SetDirPath(dirPath string) {
	this.dirPath = dirPath
}

//设置时间
func (this *LogAnalysisLogic) SetLogTime(logTime string) {
	this.logTime = logTime
}

//设置排行榜取前多少名
func (this *LogAnalysisLogic) SetHowManyInTheTop(howManyInTheTop int) {
	this.howManyInTheTop = howManyInTheTop
}

//初始化generalLogChan channel
func (this *LogAnalysisLogic) SetGeneralLogChan(total int) {
	this.generalLogChan = make(chan []string, total)
}

//
func (this *LogAnalysisLogic) SetCloseProgramChan(total int) {
	this.closeProgram = make(chan string, total)
}

//将指定路径下的zip包解压然后遍历所有文件
func (this *LogAnalysisLogic) ZipDecompress() {
	//有则跳过无则插入
	recordLogDate := time.Now().Format("2006-01-02")
	id := services.LogCollectionRecordOneServices(GeneralLogAnalysis, recordLogDate)
	if id == 0 {
		//采集日志状态入库 初始化一个未采集，采集完成更新采集状态
		services.LogCollectionRecordAddServices(this.logTime, GeneralLogAnalysis, 0, 0, NotCollected)
	}

	//获取未采集日期列表
	dateList := services.LogCollectionRecordListServices(GeneralLogAnalysis, NotCollected)
	for _, v := range dateList {
		this.logTime = v
		zipFilePath := filepath.Join(this.dirPath, this.logTime+".zip")
		_, err := os.Stat(zipFilePath)
		if err != nil {
			log.Println(this.logTime, "----没有数据包，跳过处理")
			continue
		}
		err = util.ZipHandle(zipFilePath, this.dirPath, this.logTime)
		if err != nil {
			log.Println("解压缩包出现问题", err.Error())
		}
		this.eachFile(filepath.Join(this.dirPath, this.logTime))
	}

}

//遍历解压缩包里的所有文件
func (this *LogAnalysisLogic) eachFile(eachFilePath string) {

	fileList, err := ioutil.ReadDir(eachFilePath)
	if err != nil {
		log.Println(err.Error())
	}
	i := 0
	this.taskCount = 0
	//启用协程池
	//antpool.P.Submit(func() {
	for _, v := range fileList {
		doMain := strings.Split(v.Name(), "-")[0]
		//常规日志分析报告
		if strings.Contains(v.Name(), logFileTypeList[1]) {
			//程序处理过快会导致too many open files错误，这里限制一下处理速度
			if i >= 500 {
				i = 0
				time.Sleep(time.Second * 5)
			}
			//计算处理了多少次任务
			this.taskCount++
			//获取常规日志报告数据
			filePath := filepath.Join(eachFilePath, v.Name())
			this.getFileContent(filePath, doMain)
			i++
		}
		//安全日志分析报告
		//if strings.Contains(v.Name(), logFileTypeList[2]) {
		//	//获取常规日志报告数据
		//	filePath := filepath.Join(eachFilePath, v.Name())
		//	fileContent := this.GetFileContent(filePath)
		//	if fileContent == "" {
		//		log.Println("未读取到文件内容，跳过处理", filePath)
		//		return
		//	}
		//	this.getGeneralLogData(fileContent, doMain)
		//}
	}
	//})
	handelCount := 0
	processedCount := 0 //总处理过多少个任务
	//循环读取channel数据插入到数据
	for {
		select {
		case generalLogData := <-this.generalLogChan:
			//程序处理过快会导致too many open files错误，这里限制一下处理速度
			if handelCount >= 200 {
				handelCount = 0
				time.Sleep(time.Second * 5)
			}
			if generalLogData[1] == "" {
				log.Println("未读取到文件内容，跳过处理", generalLogData[0])
				continue
			}
			//启用协程池 开启协程
			antpool.P.Submit(func() {
				this.getGeneralLogData(generalLogData[1], generalLogData[2])
			})
		case domain := <-this.closeProgram:
			//处理完成一件任务 总任务数减1
			processedCount++
			this.taskCount--
			if this.taskCount <= 0 {
				removeDirPath := filepath.Join(this.dirPath, this.logTime)
				err := os.RemoveAll(removeDirPath)
				if err != nil {
					log.Println(removeDirPath, "---删除失败")
				}
				//执行完任务 根据日期和类型更新采集状态
				services.LogCollectionRecordUpdateServices(this.logTime, GeneralLogAnalysis, this.taskCount, processedCount, Collected)
				util.LastAcquisitionDateAdd(this.logTime)
				log.Println(this.logTime, "---执行完成---本次总处理任务数为---", processedCount)
				fmt.Println(this.logTime, "程序执行完成")
				goto quit
			}
			fmt.Println(domain, "处理完成")
			log.Println(domain, "处理完成")
		}
	}

quit:
	return
}

func (this *LogAnalysisLogic) getFileContent(filePath string, doMain string) {
	generalLogData := make([]string, 3)
	fileContentByte, err := ioutil.ReadFile(filePath)
	generalLogData[0] = filePath
	generalLogData[2] = doMain
	if err != nil {
		log.Println(filePath+"function---getFileContent---文件读取失败，跳过处理", err.Error())
		generalLogData[1] = ""
	} else {
		generalLogData[1] = string(fileContentByte)
		this.generalLogChan <- generalLogData
		generalLogData = []string{}
	}
}
