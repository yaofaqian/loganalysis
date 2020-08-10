package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//死链访问分析表
type ReportDeadChain struct {
	Id             int64     `orm:"pk;auto"`                                         //自增id
	DeadSortid     int       `orm:"description(排序id)"`                               // 排序id
	DeadCount      int       `orm:"null;description(访问次数)"`                          // 访问次数
	DeadRate       float64   `orm:"null;digits(12);decimals(3);description(访问占比)"`   //访问占比
	DeadBlance     int64     `orm:"null;description(访问流量)"`                          // 访问流量
	DeadBlanceRate float64   `orm:"null;digits(12);decimals(3);description(访问流量占比)"` //访问流量占比
	DeadDomain     string    `orm:"index;description(网址)"`                           //网址
	DeadLogDate    time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	DeadMd5        string    `orm:"index;description(md5校验)"`                        //md5校验
	CreateDate     time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	DeadUrl        string    `orm:"null;description(访问页面URL)"`                       //访问页面URL
	tableName      string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportDeadChain) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportDeadChain) TableName() string {
	return "report_dead_chain202009" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportDeadChain) ReportDeadChainBatchAdd(data *[]ReportDeadChain) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
