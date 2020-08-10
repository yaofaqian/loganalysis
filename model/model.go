package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"loganalysis/config"
)

func init() {
	dbConfig := config.Config
	//tableNameSuffixMonth := time.Now().Format("2006-01")
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&loc=Local", dbConfig.DBUser, dbConfig.DBPwd, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName, dbConfig.DBCharset), 60, 60)
	orm.RegisterModel(new(ReportDataOverview), new(ReportIpTraffic), new(ReportPageView), new(ReportStaticResource), new(ReportDeadChain), new(ReportSource), new(ReportSearchCrawler), new(ReportKeyword), new(ReportRegional), new(ReportSystem), new(ReportBrowser), new(ReportStatusCode), new(LogCollectionRecord))
	orm.RunSyncdb("default", false, false) //第二个参数 是否每次都重新创建，第三个参数 是否每次都更新表结构
}
