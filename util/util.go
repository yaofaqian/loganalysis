package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/yeka/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//string数据转换为int
func DataConv(data string) (v int) {
	v, _ = strconv.Atoi(strings.ReplaceAll(data, ",", ""))
	return v
}

//string转换为float64
func DataConvFloat64(data string) (v float64) {
	v, _ = strconv.ParseFloat(data, 64)
	return v
}

//文件大小转换
func FileConv(data string) (size int64) {
	var sizeFloat float64
	uintSymbol := strings.ToLower(string(data[len(data)-1]))
	switch uintSymbol {
	case "b":
		size, _ = strconv.ParseInt(data[:len(data)-1], 10, 64)
		return size
	case "k":
		sizeFloat, _ = strconv.ParseFloat(data[:len(data)-1], 64)
		size = int64(sizeFloat * 1024)
		return size
	case "m":
		sizeFloat, _ = strconv.ParseFloat(data[:len(data)-1], 64)
		size = int64(sizeFloat * 1024 * 1024)
		return size
	case "g":
		sizeFloat, _ = strconv.ParseFloat(data[:len(data)-1], 64)
		size = int64(sizeFloat * 1024 * 1024 * 1024)
		return size
	case "t":
		sizeFloat, _ = strconv.ParseFloat(data[:len(data)-1], 64)
		size = int64(sizeFloat * 1024 * 1024 * 1024 * 1024)
		return size
	}
	return size
}

//解压zip文件方法
func ZipHandle(packagePath string, filePath string, logTime string) error {
	r, err := zip.OpenReader(packagePath)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer r.Close()
	fpath := ""
	for _, f := range r.File {
		fileName := ChineseCodeConv(f.Name)
		fpath = filepath.Join(filePath, logTime, fileName)
		//if strings.Contains(fileName, "常规分析报告") {
		//	fpath = filepath.Join(filePath, ChineseCodeConv(f.Name))
		//} else {
		//	fpath = filepath.Join(filePath, "安全分析报告.html")
		//}

		//设置解压缩密码
		f.SetPassword("_chinacourt_")
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := f.Open()
			if err != nil {
				log.Println("打开文件发生错误---", fpath, "---", err.Error())
				continue
			}
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				inFile.Close()
				log.Println("打开文件发生错误---", fpath, "---", err.Error())
				continue
			}
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				log.Println("复制文件发生错误---", fpath, "---", err.Error())
			}
			inFile.Close()
			outFile.Close()
		}
	}
	return err
}

//中文乱码处理
func ChineseCodeConv(str string) string {
	utf8Reader := transform.NewReader(strings.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	fileName, _ := ioutil.ReadAll(utf8Reader)
	s := string(fileName)
	return s
}

//md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

type LastAcquisitionDate struct {
	LastDate string `json:"last_date"`
}

//获取最后采集日期
func GetLastAcquisitionDate() (lastAcquisitionDate LastAcquisitionDate) {

	filePtr, err := os.Open("./last_date/last_date.json")
	if err != nil {
		log.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&lastAcquisitionDate)
	if err != nil {
		log.Println("Decoder failed", err.Error())

	}
	return lastAcquisitionDate
}
func LastAcquisitionDateAdd(logTime string) {
	lastAcquisitionDate := LastAcquisitionDate{logTime}

	// 创建文件
	filePtr, err := os.Create("./last_date/last_date.json")
	if err != nil {
		log.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(lastAcquisitionDate)
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

	// 带JSON缩进格式写文件
	//data, err := json.MarshalIndent(personInfo, "", "  ")
	//if err != nil {
	// fmt.Println("Encoder failed", err.Error())
	//
	//} else {
	// fmt.Println("Encoder success")
	//}
	//
	//filePtr.Write(data)
}

//获取一段时间内每天的日期
func GetDateFromRange(startTime string, endTime string) (dateList []string) {
	startTimeTmp, _ := time.Parse("2006-01-02", startTime)
	endTimeTmp, _ := time.Parse("2006-01-02", endTime)
	days := int((endTimeTmp.Unix() - startTimeTmp.Unix()) / 86400)
	for i := 1; i <= days; i++ {
		tmpTime := startTimeTmp.Unix() + int64(86400*i)
		dateList = append(dateList, time.Unix(tmpTime, 0).Format("2006-01-02"))
	}
	return dateList
}

//求并集
func UnionString(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func IntersectString(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//获取数组差集 字符串类型
func DifferenceString(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := IntersectString(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
