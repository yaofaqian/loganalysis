package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_regional表数据
func ReportRegionalBatchAddServices(regionalDataList *[]RegionData, doMain string, logTime string) {
	if len(*regionalDataList) <= 0 {
		log.Println("地域分布分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var regionalList []model.ReportRegional
	regionalData := model.ReportRegional{}
	for k, v := range *regionalDataList {
		regionalData.RegionalSortid = k
		regionalData.RegionalCount = v.RequestUserTotal
		regionalData.RegionalRate = v.RequestUserRatio
		regionalData.RegionalArea = v.CountryAndArea
		regionalData.ResourceDomain = doMain
		regionalData.ResourceLogDate, _ = time.Parse("2006-01-02", logTime)
		regionalData.ResourceMd5 = util.Md5Encryption(doMain + logTime)
		regionalList = append(regionalList, regionalData)
	}
	successNums := regionalData.ReportRegionalBatchAdd(&regionalList)
	if successNums != 0 {
		fmt.Println(doMain, "----", regionalData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
