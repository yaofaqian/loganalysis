package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_dead_chain表数据
func ReportDeadChainBatchAddServices(deadChainDataList *[]DeadChainData, doMain string, logTime string) {
	//没数据直接返回
	if len(*deadChainDataList) <= 0 {
		log.Println("死链分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var deadChainList []model.ReportDeadChain
	deadChainData := model.ReportDeadChain{}
	for _, v := range *deadChainDataList {
		deadChainData.DeadSortid = v.Id
		deadChainData.DeadCount = v.RequestTotal
		deadChainData.DeadRate = v.RequestRatio
		deadChainData.DeadBlance = v.Flow
		deadChainData.DeadBlanceRate = v.FlowRatio
		deadChainData.DeadDomain = doMain
		deadChainData.DeadLogDate, _ = time.Parse("2006-01-02", logTime)
		deadChainData.DeadMd5 = util.Md5Encryption(doMain + logTime)
		deadChainData.DeadUrl = v.Url
		deadChainList = append(deadChainList, deadChainData)
	}
	successNums := deadChainData.ReportDeadChainBatchAdd(&deadChainList)
	if successNums != 0 {
		fmt.Println(doMain, "----", deadChainData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
