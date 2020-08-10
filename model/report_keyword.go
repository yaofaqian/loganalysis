package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//搜索关键词分析表
type ReportKeyword struct {
	Id             int64     `orm:"pk;auto"`                                       //自增id
	KeywordSortid  int       `orm:"description(排序id)"`                             // 排序id
	KeywordCount   int       `orm:"null;description(访问次数)"`                        // 访问次数
	KeywordSource  string    `orm:"null;description(访问占比)"`                        //访问占比
	KeywordDomain  string    `orm:"index;description(网址)"`                         //网址
	KeywordLogDate time.Time `orm:"index;type(date);description(日志日期)"`            //日志日期
	KeywordMd5     string    `orm:"index;description(md5校验)"`                      //md5校验
	CreateDate     time.Time `orm:"auto_now_add;type(datetime);description(创建日期)"` //创建日期
	KeywordName    string    `orm:"null;description(关键词)"`                         //关键词
	tableName      string    `orm:"-"`                                             //表名
}

// 设置引擎为 MyISAM
func (this *ReportKeyword) TableEngine() string {
	return "MyISAM"
}

//定义表名
func (this *ReportKeyword) TableName() string {
	return "report_keyword2020" //+ time.Now().Format("200601") //this.tableName
}

//批量插入数据
func (this *ReportKeyword) ReportKeywordBatchAdd(data *[]ReportKeyword) int64 {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return successNums
}
