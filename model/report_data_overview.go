package model

import (
	"github.com/astaxie/beego/orm"
	"log"
)

//数据概况表
type ReportDataOverview struct {
	Id              int64  `orm:"pk;auto"`
	OverviewDomain  string `orm:"description(法院网址)"`
	OverviewLogDate string `orm:"description(日志日期)"`
	OverviewCount   int    `orm:"description(总访问量)"`
	OverviewPv      int    `orm:"description(PV数)"`
	OverviewUv      int    `orm:"description(UV数)"`
	OverviewIp      int    `orm:"description(独立IP数)"`
	OverviewError   int    `orm:"description(异常访问数)"`
	OverviewBlance  int64  `orm:"description(消耗流量)"`
	OverviewLogSize int64  `orm:"description(日志大小)"`
	OverviewMd5     string `orm:"description(md5校验)"`
	CreateDate      string `orm:"description(创建日期)"`
}

//插入一条数据
func (this *ReportDataOverview) ReportDataOverviewAdd(overview *ReportDataOverview) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(overview)
	if err != nil {
		log.Println("插入失败", overview, "---", err.Error())
	}
	return id
}
