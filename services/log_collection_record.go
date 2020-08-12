package services

import (
	"fmt"
	"loganalysis/model"
	"strconv"
	"time"
)

//插入到log_collection_record表数据
func LogCollectionRecordAddServices(logTime string, logType int, estimateCount int, actualCount int, acquisitionStatus int) {
	logCollectionRecord := model.LogCollectionRecord{}
	logCollectionRecord.RecordLogDate, _ = time.Parse("2006-01-02", logTime)
	logCollectionRecord.RecordLogType = logType
	logCollectionRecord.RecordLogEstimateCount = estimateCount
	logCollectionRecord.RecordLogActualCount = actualCount
	logCollectionRecord.RecordLogStatus = acquisitionStatus
	id := logCollectionRecord.LogCollectionRecordAdd(&logCollectionRecord)
	if id != 0 {
		fmt.Println(logCollectionRecord.TableName(), "---", strconv.Itoa(int(id)), "插入成功")
	}
}

//更新到log_collection_record表数据
func LogCollectionRecordUpdateServices(logTime string, logType int, estimateCount int, actualCount int, acquisitionStatus int) {
	logCollectionRecord := model.LogCollectionRecord{}
	logCollectionRecord.RecordLogDate, _ = time.Parse("2006-01-02", logTime)
	logCollectionRecord.RecordLogType = logType
	logCollectionRecord.RecordLogEstimateCount = estimateCount
	logCollectionRecord.RecordLogActualCount = actualCount
	logCollectionRecord.RecordLogStatus = acquisitionStatus
	logCollectionRecord.LogCollectionRecordUpdate(&logCollectionRecord)
}

//查询采集日志日期列表
func LogCollectionRecordListServices(recordLogType int, recordLogStatus int) (dateList []string) {
	logCollectionRecord := new(model.LogCollectionRecord)
	logCollectionRecord.LogCollectionRecordList(recordLogType, recordLogStatus, &dateList)
	return dateList
}

//根据recordLogType和recordLogDate查询信息
func LogCollectionRecordOneServices(recordLogType int, recordLogDate string) int64 {
	logCollectionRecord := new(model.LogCollectionRecord)
	id := logCollectionRecord.LogCollectionRecordOne(recordLogType, recordLogDate)
	return id
}
