package model

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//采集日志表
type LogCollectionRecord struct {
	Id                     int64     `orm:"pk;auto"`                                         //自增id
	RecordLogType          int       `orm:"description(采集日志类型:1 常规日志分析 2 安全日志分析)"`           // 采集日志类型
	RecordLogEstimateCount int       `orm:"null;description(预计采集数量)"`                        // 预计采集数量
	RecordLogActualCount   int       `orm:"null;description(实际采集数量)"`                        // 实际采集数量
	RecordLogDate          time.Time `orm:"index;type(date);description(采集日期)"`              //日志日期
	CreateDate             time.Time `orm:"auto_now_add;type(datetime);description(采集完成时间)"` //采集完成时间
	tableName              string    `orm:"-"`                                               //表名
}

// 设置引擎为 MyISAM
func (this *LogCollectionRecord) TableEngine() string {
	return "InnoDB"
}

//定义表名
func (this *LogCollectionRecord) TableName() string {
	return "log_collection_record" //+ time.Now().Format("200601") //this.tableName
}

//插入数据
func (this *LogCollectionRecord) LogCollectionRecordAdd(data *LogCollectionRecord) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err != nil {
		log.Println("插入失败", data, "---", err.Error())
	}
	return id
}
