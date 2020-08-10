package services

import (
	"fmt"
	"loganalysis/model"
	"strconv"
	"time"
)

//插入到log_collection_record表数据
func LogCollectionRecordAddServices(logTime string, logType int, estimateCount int, actualCount int) {
	logCollectionRecord := model.LogCollectionRecord{}
	logCollectionRecord.RecordLogDate, _ = time.Parse("2006-01-02", logTime)
	logCollectionRecord.RecordLogType = logType
	logCollectionRecord.RecordLogEstimateCount = estimateCount
	logCollectionRecord.RecordLogActualCount = actualCount
	id := logCollectionRecord.LogCollectionRecordAdd(&logCollectionRecord)
	if id != 0 {
		fmt.Println(logCollectionRecord.TableName(), "---", strconv.Itoa(int(id)), "插入成功")
	}
}
