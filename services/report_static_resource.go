package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_static_resource表数据
func ReportStaticResourceBatchAddServices(staticResourceDataList *[]StaticResourceData, doMain string, logTime string) {
	if len(*staticResourceDataList) <= 0 {
		log.Println("静态资源访问分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var staticResourceList []model.ReportStaticResource
	staticResourceData := model.ReportStaticResource{}
	for _, v := range *staticResourceDataList {
		staticResourceData.ResourceSortid = v.Id
		staticResourceData.ResourceCount = v.RequestTotal
		staticResourceData.ResourceRate = v.RequestRatio
		staticResourceData.ResourceBlance = v.Flow
		staticResourceData.ResourceBlanceRate = v.FlowRatio
		staticResourceData.ResourceDomain = doMain
		staticResourceData.ResourceLogDate, _ = time.Parse("2006-01-02", logTime)
		staticResourceData.ResourceMd5 = util.Md5Encryption(doMain + logTime)
		staticResourceData.ResourceUrl = v.Url
		staticResourceList = append(staticResourceList, staticResourceData)
	}
	successNums := staticResourceData.ReportStaticResourceBatchAdd(&staticResourceList)
	if successNums != 0 {
		fmt.Println(doMain, "----", staticResourceData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
