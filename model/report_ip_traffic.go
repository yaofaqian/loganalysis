package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//IP流量分析表
type ReportIpTraffic struct {
	Id           int64     `orm:"pk;auto"`                                         //自增id
	IpSortid     int       `orm:"description(排序id)"`                               // 排序id
	IpCount      int       `orm:"null;description(访问次数)"`                          // 访问次数
	IpRate       float64   `orm:"null;digits(12);decimals(3);description(访问占比)"`   //访问占比
	IpArea       string    `orm:"null;description(国家和地区)"`                         //国家和地区
	IpBlance     int64     `orm:"null;description(访问流量)"`                          // 访问流量
	IpBlanceRate float64   `orm:"null;digits(12);decimals(3);description(访问流量占比)"` //访问流量占比
	IpDomain     string    `orm:"index;description(网址)"`                           //网址
	IpLogDate    time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	Ip           string    `orm:"null;description(IP地址)"`                          //IP地址
	IpMd5        string    `orm:"index;description(md5校验)"`                        //md5校验
	CreateDate   time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	tableName    string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportIpTraffic) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportIpTraffic) TableName() string {
	return "report_ip_traffic202009" //+ time.Now().Format("200601") //this.tableName
}

//添加一条数据
func (this *ReportIpTraffic) ReportIpTrafficAdd(data *ReportIpTraffic) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return id
}

//批量插入数据
func (this *ReportIpTraffic) ReportIpTrafficBatchAdd(data *[]ReportIpTraffic) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
