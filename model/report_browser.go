package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//浏览器分析表
type ReportBrowser struct {
	Id              int64     `orm:"pk;auto"` //自增id
	BrowserSortid   int       // 排序id
	BrowserCount    int       `orm:"null"`                        // 访问用户
	BrowserUserRate float64   `orm:"null;digits(12);decimals(3)"` //访问用户占比
	BrowserName     string    `orm:"null"`                        //浏览器名称
	BrowserRate     float64   `orm:"null;digits(12);decimals(3)"` //浏览器占比
	BrowserDomain   string    `orm:"index"`                       //网址
	BrowserMd5      string    `orm:"index"`                       //md5校验
	BrowserLogDate  time.Time `orm:"index;type(date)"`            //日志日期
	CreateDate      time.Time `orm:"auto_now_add;type(datetime)"` //创建日期
	tableName       string    `orm:"-"`                           //表名
}

// 设置引擎为 MyISAM
func (this *ReportBrowser) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportBrowser) TableName() string {
	return "report_browser2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportBrowser) ReportBrowserBatchAdd(data *[]ReportBrowser) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
