package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_system表数据
func ReportSystemBatchAddServices(regionalDataList *[]OperatingSystemData, doMain string, logTime string) {
	if len(*regionalDataList) <= 0 {
		log.Println("操作系统分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var systemList []model.ReportSystem
	systemData := model.ReportSystem{}
	for k, v := range *regionalDataList {
		systemData.SystemSortid = k
		systemData.SystemCount = v.RequestUserTotal
		systemData.SystemUserRate = v.RequestUserRatio
		systemData.SystemName = v.OperatingSystem
		systemData.SystemRate = v.OperatingSystemRatio
		systemData.SystemDomain = doMain
		systemData.SystemLogDate, _ = time.Parse("2006-01-02", logTime)
		systemData.SystemMd5 = util.Md5Encryption(doMain + logTime)
		systemList = append(systemList, systemData)
	}
	successNums := systemData.ReportSystemBatchAdd(&systemList)
	if successNums != 0 {
		fmt.Println(doMain, "----", systemData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
