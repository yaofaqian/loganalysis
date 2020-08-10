package logic

//数据概况匹配规则
var dataOverviewRules = map[string]string{
	"RequestTotal": "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(1)",
	"Pv":           "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(2)",
	"Uv":           "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(3)",
	"Ip":           "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(4)",
	"ErrRequest":   "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(5)",
	"Flow":         "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(6)",
	"LogSize":      "#main > div > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(7)",
}

//ip流量分析匹配规则
var analysisIpTrafficRules = map[string]string{
	"PublicRule":     "#main > div > table:nth-child(8) > tbody > tr",
	"Id":             "td:nth-child(1)",
	"RequestTotal":   "td:nth-child(2)",
	"RequestRatio":   "td:nth-child(3)",
	"Ip":             "td:nth-child(4)",
	"CountryAndArea": "td:nth-child(5)",
	"Flow":           "td:nth-child(6)",
	"FlowRatio":      "td:nth-child(7)",
}

//页面访问分析匹配规则
var pageAccessRules = map[string]string{
	"PublicRule":   "#main > div > table:nth-child(11) > tbody > tr",
	"Id":           "td:nth-child(1)",
	"RequestTotal": "td:nth-child(2)",
	"RequestRatio": "td:nth-child(3)",
	"Flow":         "td:nth-child(4)",
	"FlowRatio":    "td:nth-child(5)",
	"Url":          "td:nth-child(6)",
}

//静态资源访问分析匹配规则
var staticResourcesRules = map[string]string{
	"PublicRule":   "#main > div > table:nth-child(14) > tbody > tr",
	"Id":           "td:nth-child(1)",
	"RequestTotal": "td:nth-child(2)",
	"RequestRatio": "td:nth-child(3)",
	"Flow":         "td:nth-child(4)",
	"FlowRatio":    "td:nth-child(5)",
	"Url":          "td:nth-child(6)",
}

//死链分析
var deadChainRules = map[string]string{
	"PublicRule":   "#main > div > table:nth-child(17) > tbody > tr",
	"Id":           "td:nth-child(1)",
	"RequestTotal": "td:nth-child(2)",
	"RequestRatio": "td:nth-child(3)",
	"Flow":         "td:nth-child(4)",
	"FlowRatio":    "td:nth-child(5)",
	"Url":          "td:nth-child(6)",
}

//来源分析
var sourceRules = map[string]string{
	"PublicRule":   "#main > div > table:nth-child(20) > tbody > tr",
	"Id":           "td:nth-child(1)",
	"RequestTotal": "td:nth-child(2)",
	"RequestRatio": "td:nth-child(3)",
	"Url":          "td:nth-child(6)",
}

//搜索引擎爬虫分析
var searchCrawlerRules = map[string]string{
	"PublicRule":    "#main > div > table:nth-child(23) > tbody > tr",
	"RequestTotal":  "td:nth-child(1)",
	"RequestRatio":  "td:nth-child(2)",
	"SearchCrawler": "td:nth-child(3)",
	"Flow":          "td:nth-child(4)",
	"FlowRatio":     "td:nth-child(5)",
}

//关键词分析
var keywordRules = map[string]string{
	"PublicRule":   "#main > div > table:nth-child(26) > tbody > tr",
	"Id":           "td:nth-child(1)",
	"Total":        "td:nth-child(2)",
	"Keyword":      "td:nth-child(3)",
	"SearchSource": "td:nth-child(4)",
}

//地域分布
var regionRules = map[string]string{
	"PublicRule":       "#main > div > table:nth-child(29) > tbody > tr",
	"RequestUserTotal": "td:nth-child(1)",
	"RequestUserRatio": "td:nth-child(2)",
	"CountryAndArea":   "td:nth-child(3)",
}

//操作系统
var operatingSystemRules = map[string]string{
	"PublicRule":               "#main > div > table:nth-child(32) > tbody > tr",
	"RequestUserTotal":         "td:nth-child(1)",
	"RequestUserRatio":         "td:nth-child(2)",
	"OperatingSystem":          "td:nth-child(3)",
	"OperatingSystemRatio":     "td.graph > div",
	"OperatingSystemRatioAttr": "style",
}

//浏览器
var browserRules = map[string]string{
	"PublicRule":       "#main > div > table:nth-child(35) > tbody > tr",
	"RequestUserTotal": "td:nth-child(1)",
	"RequestUserRatio": "td:nth-child(2)",
	"BrowserName":      "td:nth-child(3)",
	"BrowserRatio":     "td.graph > div",
	"BrowserRatioAttr": "style",
}

//状态码
var statusCodeRules = map[string]string{
	"PublicRule": "#main > div > table:nth-child(38) > tbody > tr",
	"Total":      "td:nth-child(1)",
	"Ratio":      "td:nth-child(2)",
	"StatusCode": "td:nth-child(3)",
}
