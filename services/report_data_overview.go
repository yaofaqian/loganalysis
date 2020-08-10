package services

import (
	"fmt"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//插入到report_data_overview表一条数据
func ReportDataOverviewAddServices(dataOverview *DataOverview, doMain string, logTime string) {
	overview := new(model.ReportDataOverview)
	overview.OverviewCount = dataOverview.RequestTotal
	overview.OverviewPv = dataOverview.PvTotal
	overview.OverviewUv = dataOverview.UvTotal
	overview.OverviewBlance = dataOverview.Flow
	overview.OverviewError = dataOverview.ErrorRequestTotal
	overview.OverviewIp = dataOverview.IpTotal
	overview.OverviewDomain = doMain
	overview.OverviewLogDate = logTime
	overview.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	overview.OverviewMd5 = util.Md5Encryption(doMain + logTime)
	overview.OverviewLogSize = dataOverview.LogSize
	id := overview.ReportDataOverviewAdd(overview)
	if id != 0 {
		fmt.Println("report_data_overview table ---", strconv.Itoa(int(id)), "插入成功")
	}
}
