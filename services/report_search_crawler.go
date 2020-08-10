package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_search_crawler表数据
func ReportSearchCrawlerBatchAddServices(searchCrawlerDataList *[]SearchCrawlerData, doMain string, logTime string) {
	if len(*searchCrawlerDataList) <= 0 {
		log.Println("搜索引擎爬虫分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var searchCrawlerList []model.ReportSearchCrawler
	staticResourceData := model.ReportSearchCrawler{}
	for k, v := range *searchCrawlerDataList {
		staticResourceData.CrawlerSortid = k
		staticResourceData.CrawlerCount = v.RequestTotal
		staticResourceData.CrawlerRate = v.RequestRatio
		staticResourceData.CrawlerBlance = v.Flow
		staticResourceData.CrawlerBlanceRate = v.FlowRatio
		staticResourceData.CrawlerDomain = doMain
		staticResourceData.CrawlerLogDate, _ = time.Parse("2006-01-02", logTime)
		staticResourceData.CrawlerMd5 = util.Md5Encryption(doMain + logTime)
		staticResourceData.CrawlerName = v.SearchCrawler
		searchCrawlerList = append(searchCrawlerList, staticResourceData)
	}
	successNums := staticResourceData.ReportSearchCrawlerBatchAdd(&searchCrawlerList)
	if successNums != 0 {
		fmt.Println(doMain, "----", staticResourceData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
