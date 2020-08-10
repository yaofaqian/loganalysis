package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//状态码分析表
type ReportStatusCode struct {
	Id              int64     `orm:"pk;auto"`                                       //自增id
	CodeSortid      int       `orm:"description(排序id)"`                             // 排序id
	CodeCount       int       `orm:"null;description(次数)"`                          // 次数
	CodeRate        float64   `orm:"null;digits(12);decimals(3);description(占比)"`   //占比
	CodeDescription string    `orm:"null;description(状态码)"`                         //状态码
	CodeDomain      string    `orm:"index;description(网址)"`                         //网址
	CodeMd5         string    `orm:"index;description(md5校验)"`                      //md5校验
	CodeLogDate     time.Time `orm:"index;type(date);description(日志日期)"`            //日志日期
	CreateDate      time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"` //创建日期
	tableName       string    `orm:"-"`                                             //表名
}

// 设置引擎为 MyISAM
func (this *ReportStatusCode) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportStatusCode) TableName() string {
	return "report_status_code2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportStatusCode) ReportStatusCodeBatchAdd(data *[]ReportStatusCode) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
