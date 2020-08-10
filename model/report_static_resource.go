package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//静态资源访问分析表
type ReportStaticResource struct {
	Id                 int64     `orm:"pk;auto"`                                         //自增id
	ResourceSortid     int       `orm:"description(排序id)"`                               // 排序id
	ResourceCount      int       `orm:"null;description(访问次数)"`                          // 访问次数
	ResourceRate       float64   `orm:"null;digits(12);decimals(3);description(访问占比)"`   //访问占比
	ResourceBlance     int64     `orm:"null;description(访问流量)"`                          // 访问流量
	ResourceBlanceRate float64   `orm:"null;digits(12);decimals(3);description(访问流量占比)"` //访问流量占比
	ResourceDomain     string    `orm:"index;description(网址)"`                           //网址
	ResourceLogDate    time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	ResourceMd5        string    `orm:"index;description(md5校验)"`                        //md5校验
	CreateDate         time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	ResourceUrl        string    `orm:"null;description(访问页面URL)"`                       //访问页面URL
	tableName          string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportStaticResource) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportStaticResource) TableName() string {
	return "report_static_resource202009" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportStaticResource) ReportStaticResourceBatchAdd(data *[]ReportStaticResource) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
