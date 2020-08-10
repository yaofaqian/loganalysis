package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//搜索引擎爬虫分析表
type ReportSearchCrawler struct {
	Id                int64     `orm:"pk;auto"`                                         //自增id
	CrawlerSortid     int       `orm:"description(排序id)"`                               // 排序id
	CrawlerCount      int       `orm:"null;description(访问次数)"`                          // 访问次数
	CrawlerRate       float64   `orm:"null;digits(12);decimals(3);description(访问占比)"`   //访问占比
	CrawlerBlance     int64     `orm:"null;description(访问流量)"`                          // 访问流量
	CrawlerBlanceRate float64   `orm:"null;digits(12);decimals(3);description(访问流量占比)"` //访问流量占比
	CrawlerDomain     string    `orm:"index;description(网址)"`                           //网址
	CrawlerLogDate    time.Time `orm:"index;type(date);description(日志日期)"`              //日志日期
	CrawlerMd5        string    `orm:"index;description(md5校验)"`                        //md5校验
	CreateDate        time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"`   //创建日期
	CrawlerName       string    `orm:"null;description(爬虫名称)"`                          //爬虫名称
	tableName         string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *ReportSearchCrawler) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportSearchCrawler) TableName() string {
	return "report_search_crawler2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportSearchCrawler) ReportSearchCrawlerBatchAdd(data *[]ReportSearchCrawler) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
