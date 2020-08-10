package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_page_view表数据
func ReportPageViewBatchAddServices(pageVidwDataList *[]PageAccessData, doMain string, logTime string) {
	if len(*pageVidwDataList) <= 0 {
		log.Println("访问页面分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var pageViewList []model.ReportPageView
	pageViewData := model.ReportPageView{}
	for _, v := range *pageVidwDataList {
		pageViewData.ViewSortid = v.Id
		pageViewData.ViewCount = v.RequestTotal
		pageViewData.ViewRate = v.RequestRatio
		pageViewData.ViewBlance = v.Flow
		pageViewData.ViewBlanceRate = v.FlowRatio
		pageViewData.ViewDomain = doMain
		pageViewData.ViewLogDate, _ = time.Parse("2006-01-02", logTime)
		pageViewData.ViewMd5 = util.Md5Encryption(doMain + logTime)
		pageViewData.ViewUrl = v.Url
		pageViewList = append(pageViewList, pageViewData)
	}
	successNums := pageViewData.ReportPageViewBatchAdd(&pageViewList)
	if successNums != 0 {
		fmt.Println(doMain, "----", pageViewData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
