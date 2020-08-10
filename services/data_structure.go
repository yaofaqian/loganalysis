package services

//数据概况
type DataOverview struct {
	RequestTotal      int
	PvTotal           int
	UvTotal           int
	IpTotal           int
	ErrorRequestTotal int
	Flow              int64
	LogSize           int64
}

//IP流量分析
type AnalysisIpTrafficData struct {
	Id             int
	RequestTotal   int
	RequestRatio   float64
	Ip             string
	CountryAndArea string
	Flow           int64
	FlowRatio      float64
}

//页面访问分析
type PageAccessData struct {
	Id           int
	RequestTotal int
	RequestRatio float64
	Flow         int64
	FlowRatio    float64
	Url          string
}

//静态资源访问分析
type StaticResourceData struct {
	Id           int
	RequestTotal int
	RequestRatio float64
	Flow         int64
	FlowRatio    float64
	Url          string
}

//死链访问分析
type DeadChainData struct {
	Id           int
	RequestTotal int
	RequestRatio float64
	Flow         int64
	FlowRatio    float64
	Url          string
}

//来源分析
type SourceData struct {
	Id           int
	RequestTotal int
	RequestRatio float64
	Url          string
}

//搜索引擎爬虫分析
type SearchCrawlerData struct {
	RequestTotal  int
	RequestRatio  float64
	SearchCrawler string
	Flow          int64
	FlowRatio     float64
}

//关键词分析
type KeywordData struct {
	Id           int
	Total        int
	KeyWord      string
	SearchSource string
}

//地域分布分析
type RegionData struct {
	RequestUserTotal int
	RequestUserRatio float64
	CountryAndArea   string
}

//操作系统
type OperatingSystemData struct {
	RequestUserTotal     int
	RequestUserRatio     float64
	OperatingSystem      string
	OperatingSystemRatio float64
}

//浏览器
type BrowserData struct {
	RequestUserTotal int
	RequestUserRatio float64
	BrowserName      string
	BrowserRatio     float64
}

//状态码
type StatusCodeData struct {
	Total      int
	Ratio      float64
	StatusCode string
}
