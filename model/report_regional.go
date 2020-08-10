package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//地域分布分析表
type ReportRegional struct {
	Id              int64     `orm:"pk;auto"`                                         //自增id
	RegionalSortid  int       `orm:"description(排序id)"`                               // 排序id
	RegionalCount   int       `orm:"null;description(访问用户)"`                          // 访问用户
	RegionalRate    float64   `orm:"null;digits(12);decimals(3);description(访问用户占比)"` //访问用户占比
	RegionalArea    string    `orm:"null;description(国家和地区)"`                         //国家和地区
	ResourceDomain  string    `orm:"index;description(网址)"`                           //网址
	ResourceMd5     string    `orm:"index;description(md5校验)"`                        //md5校验
	ResourceLogDate time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	CreateDate      time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	tableName       string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportRegional) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportRegional) TableName() string {
	return "report_regional2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportRegional) ReportRegionalBatchAdd(data *[]ReportRegional) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
