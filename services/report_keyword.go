package services

import (
	"fmt"
	"log"
	"loganalysis/model"
	"loganalysis/util"
	"strconv"
	"time"
)

//批量插入到report_keyword表数据
func ReportKeywordBatchAddServices(keyWordDataDataList *[]KeywordData, doMain string, logTime string) {
	if len(*keyWordDataDataList) <= 0 {
		log.Println("关键词分析板块无数据，跳过处理---", doMain, "---", logTime)
		return
	}
	var keywordList []model.ReportKeyword
	keywordData := model.ReportKeyword{}
	for _, v := range *keyWordDataDataList {
		keywordData.KeywordSortid = v.Id
		keywordData.KeywordCount = v.Total
		keywordData.KeywordName = v.KeyWord
		keywordData.KeywordSource = v.SearchSource
		keywordData.KeywordDomain = doMain
		keywordData.KeywordLogDate, _ = time.Parse("2006-01-02", logTime)
		keywordData.KeywordMd5 = util.Md5Encryption(doMain + logTime)
		keywordList = append(keywordList, keywordData)
	}
	successNums := keywordData.ReportKeywordBatchAdd(&keywordList)
	if successNums != 0 {
		fmt.Println(doMain, "----", keywordData.TableName(), "---", strconv.Itoa(int(successNums)), "插入成功")
	}
}
