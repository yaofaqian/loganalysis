package util

import (
	"crypto/md5"
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
		fmt.Println(err.Error())
		return err
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
