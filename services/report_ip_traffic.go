package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//插入到report_ip_traffic表一条数据
func ReportIpTrafficAddServices(data *AnalysisIpTrafficData, doMain string, logTime string) {
	trafficData := new(model.ReportIpTraffic)
	trafficData.IpSortid = data.Id
	trafficData.IpCount = data.RequestTotal
	trafficData.IpRate = data.RequestRatio
	trafficData.IpArea = data.CountryAndArea
	trafficData.IpBlance = data.Flow
	trafficData.IpBlanceRate = data.FlowRatio
	trafficData.IpDomain = doMain
	trafficData.IpLogDate, _ = time.Parse("2006-01-02", logTime)
	//trafficData.CreateDate = time.Now()
	trafficData.Ip = data.Ip
	trafficData.IpMd5 = util.Md5Encryption(doMain + logTime)
	id := trafficData.ReportIpTrafficAdd(trafficData)
	if id != 0 {
		fmt.Println(trafficData.TableName(), "---", strconv.Itoa(int(id)), "插入成功")
	}
}

//批量插入到report_ip_traffic表数据
func ReportIpTrafficBatchAddServices(trafficDataList *[]AnalysisIpTrafficData, doMain string, logTime string) {
	//没数据直接返回
	if len(*trafficDataList) <= 0 {
		log.Println("IP流量分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var trafficList []model.ReportIpTraffic
	trafficData := model.ReportIpTraffic{}
	for _, v := range *trafficDataList {
		trafficData.IpSortid = v.Id
		trafficData.IpCount = v.RequestTotal
		trafficData.IpRate = v.RequestRatio
		trafficData.IpArea = v.CountryAndArea
		trafficData.IpBlance = v.Flow
		trafficData.IpBlanceRate = v.FlowRatio
		trafficData.IpDomain = doMain
		trafficData.IpLogDate, _ = time.Parse("2006-01-02", logTime)
		trafficData.Ip = v.Ip
		trafficData.IpMd5 = util.Md5Encryption(doMain + logTime)
		trafficList = append(trafficList, trafficData)
	}
	successNums := trafficData.ReportIpTrafficBatchAdd(&trafficList)
	if successNums != 0 {
		fmt.Println(doMain, "----", trafficData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
