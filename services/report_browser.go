package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_browser表数据
func ReportBrowserBatchAddServices(regionalDataList *[]BrowserData, doMain string, logTime string) {
	//没数据直接返回
	if len(*regionalDataList) <= 0 {
		log.Println("浏览器分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var browserList []model.ReportBrowser
	browserData := model.ReportBrowser{}
	for k, v := range *regionalDataList {
		browserData.BrowserSortid = k
		browserData.BrowserCount = v.RequestUserTotal
		browserData.BrowserUserRate = v.RequestUserRatio
		browserData.BrowserName = v.BrowserName
		browserData.BrowserRate = v.BrowserRatio
		browserData.BrowserDomain = doMain
		browserData.BrowserLogDate, _ = time.Parse("2006-01-02", logTime)
		browserData.BrowserMd5 = util.Md5Encryption(doMain + logTime)
		browserList = append(browserList, browserData)
	}
	successNums := browserData.ReportBrowserBatchAdd(&browserList)
	if successNums != 0 {
		fmt.Println(doMain, "----", browserData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
