package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//操作系统分析表
type ReportSystem struct {
	Id             int64     `orm:"pk;auto"`                                         //自增id
	SystemSortid   int       `orm:"description(排序id)"`                               // 排序id
	SystemCount    int       `orm:"null;description(访问用户)"`                          // 访问用户
	SystemUserRate float64   `orm:"null;digits(12);decimals(3);description(访问用户占比)"` //访问用户占比
	SystemName     string    `orm:"null;description(操作系统名称)"`                        //操作系统名称
	SystemRate     float64   `orm:"null;digits(12);decimals(3);description(操作系统占比)"` //操作系统占比
	SystemDomain   string    `orm:"index;description(网址)"`                           //网址
	SystemMd5      string    `orm:"index;description(md5校验)"`                        //md5校验
	SystemLogDate  time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	CreateDate     time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	tableName      string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportSystem) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportSystem) TableName() string {
	return "report_system2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportSystem) ReportSystemBatchAdd(data *[]ReportSystem) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
