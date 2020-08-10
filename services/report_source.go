package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_source表数据
func ReportSourceBatchAddServices(sourceDataList *[]SourceData, doMain string, logTime string) {
	if len(*sourceDataList) <= 0 {
		log.Println("来源分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var sourceList []model.ReportSource
	sourceData := model.ReportSource{}
	for _, v := range *sourceDataList {
		sourceData.SourceSortid = v.Id
		sourceData.SourceCount = v.RequestTotal
		sourceData.SourceRate = v.RequestRatio
		sourceData.SourceDomain = doMain
		sourceData.SourceLogDate, _ = time.Parse("2006-01-02", logTime)
		sourceData.SourceMd5 = util.Md5Encryption(doMain + logTime)
		sourceData.SourceUrl = v.Url
		sourceList = append(sourceList, sourceData)
	}
	successNums := sourceData.ReportSourceBatchAdd(&sourceList)
	if successNums != 0 {
		fmt.Println(doMain, "----", sourceData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
