package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_status_code表数据
func ReportStatusCodeBatchAddServices(statusCodeDataList *[]StatusCodeData, doMain string, logTime string) {
	if len(*statusCodeDataList) <= 0 {
		log.Println("状态码分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var statusCodeList []model.ReportStatusCode
	statusCodeData := model.ReportStatusCode{}
	for k, v := range *statusCodeDataList {
		statusCodeData.CodeSortid = k
		statusCodeData.CodeCount = v.Total
		statusCodeData.CodeRate = v.Ratio
		statusCodeData.CodeDescription = v.StatusCode
		statusCodeData.CodeDomain = doMain
		statusCodeData.CodeLogDate, _ = time.Parse("2006-01-02", logTime)
		statusCodeData.CodeMd5 = util.Md5Encryption(doMain + logTime)
		statusCodeList = append(statusCodeList, statusCodeData)
	}
	successNums := statusCodeData.ReportStatusCodeBatchAdd(&statusCodeList)
	if successNums != 0 {
		fmt.Println(doMain, "----", statusCodeData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
