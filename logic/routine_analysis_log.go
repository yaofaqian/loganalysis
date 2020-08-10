package logic

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"loganalysis/services"
	"loganalysis/util"
	"os"
	"path/filepath"
	"strings"
)

//匹配常规分析报告数据
func (this *LogAnalysisLogic) getGeneralLogData(filContent string, doMain string) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(filContent))
	if err != nil {
		log.Println("function---getGeneralLogData---文件读取失败，跳过处理", err.Error())
		return
	}
	//匹配数据概况 插入到数据库
	dataOverview := services.DataOverview{}
	this.getDataOverviewData(document, &dataOverview)
	services.ReportDataOverviewAddServices(&dataOverview, doMain, this.logTime)
	//匹配IP流量分析
	analysisIpTrafficDataList := this.getAnalysisIpTrafficData(document)
	services.ReportIpTrafficBatchAddServices(&analysisIpTrafficDataList, doMain, this.logTime)
	//匹配页面访问分析
	pageAccessDataList := this.getPageAccessData(document)
	services.ReportPageViewBatchAddServices(&pageAccessDataList, doMain, this.logTime)
	//匹配静态资源访问分析
	staticResourceDataList := this.getStaticResourceData(document)
	services.ReportStaticResourceBatchAddServices(&staticResourceDataList, doMain, this.logTime)
	//匹配死链访问分析
	deadChainDataList := this.getDeadChainData(document)
	services.ReportDeadChainBatchAddServices(&deadChainDataList, doMain, this.logTime)
	//匹配来源分析
	sourceDataList := this.getSourceData(document)
	services.ReportSourceBatchAddServices(&sourceDataList, doMain, this.logTime)
	//匹配搜索引擎爬虫分析
	searchCrawlerDataList := this.getSearchCrawlerData(document)
	services.ReportSearchCrawlerBatchAddServices(&searchCrawlerDataList, doMain, this.logTime)
	//匹配搜索关键词分析
	keyWordDataList := this.getKeywordData(document)
	services.ReportKeywordBatchAddServices(&keyWordDataList, doMain, this.logTime)
	//匹配地域分布分析
	regionDataList := this.getRegionData(document)
	services.ReportRegionalBatchAddServices(&regionDataList, doMain, this.logTime)
	//匹配操作系统分析
	operatingSystemDataList := this.getOperatingSystemData(document)
	services.ReportSystemBatchAddServices(&operatingSystemDataList, doMain, this.logTime)
	//匹配浏览器分析
	browserDataList := this.getBrowserData(document)
	services.ReportBrowserBatchAddServices(&browserDataList, doMain, this.logTime)
	//匹配状态码分析
	statusCodeDataList := this.getStatusCodeData(document)
	services.ReportStatusCodeBatchAddServices(&statusCodeDataList, doMain, this.logTime)
	//处理完成删除解压的文件减少占用的资源空间
	go func() {
		removeFilePath := filepath.Join(this.dirPath, this.logTime, doMain+"-"+logFileTypeList[1]+".html")
		err = os.Remove(removeFilePath)
		if err != nil {
			log.Println(removeFilePath, "---删除失败")
		}
	}()
	go func() { this.closeProgram <- doMain }()
}

//匹配数据概况
func (this *LogAnalysisLogic) getDataOverviewData(document *goquery.Document, dataOverview *services.DataOverview) {
	DataOverviewMap := map[string]string{}
	totalString := ""
	for k, v := range dataOverviewRules {
		totalString = document.Find(v).Text()
		DataOverviewMap[k] = strings.ReplaceAll(totalString, " ", "")
	}
	dataOverview.RequestTotal = util.DataConv(DataOverviewMap["RequestTotal"])    //总访问量
	dataOverview.PvTotal = util.DataConv(DataOverviewMap["Pv"])                   //PV量
	dataOverview.UvTotal = util.DataConv(DataOverviewMap["Uv"])                   //UV量
	dataOverview.IpTotal = util.DataConv(DataOverviewMap["Ip"])                   // 独立IP数
	dataOverview.ErrorRequestTotal = util.DataConv(DataOverviewMap["ErrRequest"]) //异常访问数
	dataOverview.Flow = util.FileConv(DataOverviewMap["Flow"])                    //消耗流量
	dataOverview.LogSize = util.FileConv(DataOverviewMap["LogSize"])              //日志文件大小
}

//匹配IP流量分析
func (this *LogAnalysisLogic) getAnalysisIpTrafficData(document *goquery.Document) (analysisIpTrafficDataList []services.AnalysisIpTrafficData) {
	var analysisIpTrafficData services.AnalysisIpTrafficData
	document.Find(analysisIpTrafficRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		if i > this.howManyInTheTop {
			return
		}
		analysisIpTrafficData = services.AnalysisIpTrafficData{}
		analysisIpTrafficData.Id = util.DataConv(selection.Find(analysisIpTrafficRules["Id"]).Text())                                                         //ID
		analysisIpTrafficData.RequestTotal = util.DataConv(selection.Find(analysisIpTrafficRules["RequestTotal"]).Text())                                     //访问次数
		analysisIpTrafficData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(analysisIpTrafficRules["RequestRatio"]).Text(), "%", "")) //访问占比
		analysisIpTrafficData.Ip = selection.Find(analysisIpTrafficRules["Ip"]).Text()                                                                        //访问ip
		analysisIpTrafficData.CountryAndArea = selection.Find(analysisIpTrafficRules["CountryAndArea"]).Text()
		FlowString := selection.Find(analysisIpTrafficRules["Flow"]).Text()
		if FlowString != "" {
			analysisIpTrafficData.Flow = util.FileConv(strings.ReplaceAll(FlowString, " ", "")) //流量
		}
		analysisIpTrafficData.FlowRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(analysisIpTrafficRules["FlowRatio"]).Text(), "%", "")) //流量占比
		analysisIpTrafficDataList = append(analysisIpTrafficDataList, analysisIpTrafficData)
	})
	return analysisIpTrafficDataList
}

//匹配页面访问分析
func (this *LogAnalysisLogic) getPageAccessData(document *goquery.Document) (pageAccessDataList []services.PageAccessData) {
	var pageAccessData services.PageAccessData

	document.Find(pageAccessRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		if i > this.howManyInTheTop {
			return
		}
		pageAccessData = services.PageAccessData{}
		pageAccessData.Id = util.DataConv(selection.Find(pageAccessRules["Id"]).Text())                                                         //ID
		pageAccessData.RequestTotal = util.DataConv(selection.Find(pageAccessRules["RequestTotal"]).Text())                                     //访问次数
		pageAccessData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(pageAccessRules["RequestRatio"]).Text(), "%", "")) //访问占比 		//访问ip
		FlowString := selection.Find(pageAccessRules["Flow"]).Text()
		if FlowString != "" {
			pageAccessData.Flow = util.FileConv(strings.ReplaceAll(FlowString, " ", "")) //流量
		}
		pageAccessData.FlowRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(pageAccessRules["FlowRatio"]).Text(), "%", "")) //流量占比
		pageAccessData.Url = selection.Find(pageAccessRules["Url"]).Text()
		pageAccessDataList = append(pageAccessDataList, pageAccessData)
	})
	return pageAccessDataList
}

//匹配静态资源访问分析
func (this *LogAnalysisLogic) getStaticResourceData(document *goquery.Document) (staticResourceDataList []services.StaticResourceData) {
	var staticResourceData services.StaticResourceData

	document.Find(staticResourcesRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		if i > this.howManyInTheTop {
			return
		}
		staticResourceData = services.StaticResourceData{}
		staticResourceData.Id = util.DataConv(selection.Find(staticResourcesRules["Id"]).Text())                                                         //ID
		staticResourceData.RequestTotal = util.DataConv(selection.Find(staticResourcesRules["RequestTotal"]).Text())                                     //访问次数
		staticResourceData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(staticResourcesRules["RequestRatio"]).Text(), "%", "")) //访问占比 		//访问ip
		FlowString := selection.Find(staticResourcesRules["Flow"]).Text()
		if FlowString != "" {
			staticResourceData.Flow = util.FileConv(strings.ReplaceAll(FlowString, " ", "")) //流量
		}
		staticResourceData.FlowRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(staticResourcesRules["FlowRatio"]).Text(), "%", "")) //流量占比
		staticResourceData.Url = selection.Find(staticResourcesRules["Url"]).Text()
		staticResourceDataList = append(staticResourceDataList, staticResourceData)
	})
	return staticResourceDataList
}

//匹配死链访问分析
func (this *LogAnalysisLogic) getDeadChainData(document *goquery.Document) (deadChainDataList []services.DeadChainData) {
	var deadChainData services.DeadChainData

	document.Find(deadChainRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		if i > this.howManyInTheTop {
			return
		}
		deadChainData = services.DeadChainData{}
		deadChainData.Id = util.DataConv(selection.Find(deadChainRules["Id"]).Text())                                                         //ID
		deadChainData.RequestTotal = util.DataConv(selection.Find(deadChainRules["RequestTotal"]).Text())                                     //访问次数
		deadChainData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(deadChainRules["RequestRatio"]).Text(), "%", "")) //访问占比 		//访问ip
		FlowString := selection.Find(deadChainRules["Flow"]).Text()
		if FlowString != "" {
			deadChainData.Flow = util.FileConv(strings.ReplaceAll(FlowString, " ", "")) //流量
		}
		deadChainData.FlowRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(deadChainRules["FlowRatio"]).Text(), "%", "")) //流量占比
		deadChainData.Url = selection.Find(deadChainRules["Url"]).Text()
		deadChainDataList = append(deadChainDataList, deadChainData)
	})
	return deadChainDataList
}

//匹配来源分析
func (this *LogAnalysisLogic) getSourceData(document *goquery.Document) (sourceDataList []services.SourceData) {
	var sourceData services.SourceData

	document.Find(deadChainRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		if i > this.howManyInTheTop {
			return
		}
		sourceData = services.SourceData{}
		sourceData.Id = util.DataConv(selection.Find(deadChainRules["Id"]).Text())                                                         //ID
		sourceData.RequestTotal = util.DataConv(selection.Find(deadChainRules["RequestTotal"]).Text())                                     //访问次数
		sourceData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(deadChainRules["RequestRatio"]).Text(), "%", "")) //访问占比 		//访问ip
		sourceData.Url = selection.Find(deadChainRules["Url"]).Text()
		sourceDataList = append(sourceDataList, sourceData)
	})
	return sourceDataList
}

//匹配搜索引擎爬虫分析
func (this *LogAnalysisLogic) getSearchCrawlerData(document *goquery.Document) (searchCrawlerDataList []services.SearchCrawlerData) {
	var searchCrawlerData services.SearchCrawlerData

	document.Find(searchCrawlerRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		searchCrawlerData = services.SearchCrawlerData{}
		searchCrawlerData.RequestTotal = util.DataConv(selection.Find(searchCrawlerRules["RequestTotal"]).Text())                                     //访问次数
		searchCrawlerData.RequestRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(searchCrawlerRules["RequestRatio"]).Text(), "%", "")) //访问占比
		searchCrawlerData.SearchCrawler = selection.Find(searchCrawlerRules["SearchCrawler"]).Text()
		FlowString := selection.Find(searchCrawlerRules["Flow"]).Text()
		if FlowString != "" {
			searchCrawlerData.Flow = util.FileConv(strings.ReplaceAll(FlowString, " ", "")) //流量
		}
		searchCrawlerData.FlowRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(searchCrawlerRules["FlowRatio"]).Text(), "%", "")) //流量占比

		searchCrawlerDataList = append(searchCrawlerDataList, searchCrawlerData)
	})
	return searchCrawlerDataList
}

//匹配搜索关键词分析
func (this *LogAnalysisLogic) getKeywordData(document *goquery.Document) (keyWordDataList []services.KeywordData) {
	var keywordData services.KeywordData

	document.Find(keywordRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		keywordData = services.KeywordData{}
		keywordData.Id = util.DataConv(selection.Find(deadChainRules["Id"]).Text())     //ID
		keywordData.Total = util.DataConv(selection.Find(keywordRules["Total"]).Text()) //次数
		keywordData.KeyWord = selection.Find(keywordRules["KeyWord"]).Text()            //关键词
		keywordData.SearchSource = selection.Find(keywordRules["SearchSource"]).Text()  //来源
		keyWordDataList = append(keyWordDataList, keywordData)
	})
	return keyWordDataList
}

//匹配地域分布分析
func (this *LogAnalysisLogic) getRegionData(document *goquery.Document) (regionDataList []services.RegionData) {
	var regionData services.RegionData

	document.Find(regionRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		regionData = services.RegionData{}
		regionData.RequestUserTotal = util.DataConv(selection.Find(regionRules["RequestUserTotal"]).Text())                                     //访问用户数
		regionData.RequestUserRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(regionRules["RequestUserRatio"]).Text(), "%", "")) //访问用户数占比
		regionData.CountryAndArea = selection.Find(regionRules["CountryAndArea"]).Text()                                                        //国家和地区
		regionDataList = append(regionDataList, regionData)
	})
	return regionDataList
}

//匹配操作系统分析
func (this *LogAnalysisLogic) getOperatingSystemData(document *goquery.Document) (operatingSystemDataList []services.OperatingSystemData) {
	var operatingSystemData services.OperatingSystemData

	document.Find(operatingSystemRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		operatingSystemData = services.OperatingSystemData{}
		operatingSystemData.RequestUserTotal = util.DataConv(selection.Find(operatingSystemRules["RequestUserTotal"]).Text())                                     //访问用户数
		operatingSystemData.RequestUserRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(operatingSystemRules["RequestUserRatio"]).Text(), "%", "")) //访问用户数占比
		operatingSystemData.OperatingSystem = selection.Find(operatingSystemRules["OperatingSystem"]).Text()                                                      //操作系统
		operatingSystemRatio, _ := selection.Find(operatingSystemRules["OperatingSystemRatio"]).Attr(operatingSystemRules["OperatingSystemRatioAttr"])
		operatingSystemData.OperatingSystemRatio = util.DataConvFloat64(strings.ReplaceAll(strings.ReplaceAll(operatingSystemRatio, "width: ", ""), "%", "")) //操作系统占比
		operatingSystemDataList = append(operatingSystemDataList, operatingSystemData)
	})

	return operatingSystemDataList
}

//匹配浏览器分析
func (this *LogAnalysisLogic) getBrowserData(document *goquery.Document) (browserDataList []services.BrowserData) {
	var browserData services.BrowserData

	document.Find(browserRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		browserData = services.BrowserData{}
		browserData.RequestUserTotal = util.DataConv(selection.Find(browserRules["RequestUserTotal"]).Text())                                     //访问用户数
		browserData.RequestUserRatio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(browserRules["RequestUserRatio"]).Text(), "%", "")) //访问用户数占比
		browserData.BrowserName = selection.Find(browserRules["BrowserName"]).Text()                                                              //浏览器
		browserRatio, _ := selection.Find(browserRules["BrowserRatio"]).Attr(browserRules["BrowserRatioAttr"])
		browserData.BrowserRatio = util.DataConvFloat64(strings.ReplaceAll(strings.ReplaceAll(browserRatio, "width: ", ""), "%", "")) //浏览器占比
		browserDataList = append(browserDataList, browserData)
	})

	return browserDataList
}

//匹配状态码分析
func (this *LogAnalysisLogic) getStatusCodeData(document *goquery.Document) (statusCodeDataList []services.StatusCodeData) {
	var statusCodeData services.StatusCodeData

	document.Find(statusCodeRules["PublicRule"]).Each(func(i int, selection *goquery.Selection) {
		statusCodeData = services.StatusCodeData{}
		statusCodeData.Total = util.DataConv(selection.Find(statusCodeRules["Total"]).Text())                                     //次数
		statusCodeData.Ratio = util.DataConvFloat64(strings.ReplaceAll(selection.Find(statusCodeRules["Ratio"]).Text(), "%", "")) //占比
		statusCodeData.StatusCode = selection.Find(statusCodeRules["StatusCode"]).Text()                                          //状态码
		statusCodeDataList = append(statusCodeDataList, statusCodeData)
	})

	return statusCodeDataList
}
