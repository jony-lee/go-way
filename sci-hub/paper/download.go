package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//type Paper struct {
//	Doi string
//	Name string
//	BaseUrl string
//	Url string
//
//}

func Download(doi string) (filename string, err error) {

	defer func() {
		if recover() != nil {
			fmt.Println("下载失败，不支持的doi")
		}
	}()
	//doi校验
	if strings.Split(doi, ".")[0] != "10" || len(strings.Split(doi, "/")) != 2 {
		return "", errors.New("doi格式错误")
	}
	//获取真实下载地址
	var BaseUrl = "http://sci-hub.se/"
	var url = BaseUrl + doi
	content, _, err := parseUrl(url)
	re := regexp.MustCompile("href='(.*?)\\?download=true")
	downloadUrl := re.FindStringSubmatch(string(content))[1]
	if downloadUrl[0] == '/' {
		downloadUrl = "http:" + downloadUrl
	}
	//获取下载地址的content
	filename = filepath.Base(downloadUrl)
	content, header, err := parseUrl(downloadUrl)
	if err != nil {
		return
	}
	//通过Content-Type判断是pdf还是验证码页面
	var contentType = header.Get("Content-Type")
	if contentType != "application/octet-stream" {
		return "", errors.New("当前文献需要验证码")
	}
	//校验文档字节长度
	var contentLength = header.Get("Content-Length")
	length, _ := strconv.ParseInt(contentLength, 0, 32)
	if len(content) != int(length) {
		return "", errors.New("文档长度错误！")
	}
	//判断存储路径，没有则创建
	_, err = os.Stat("./pdf")
	if err != nil {
		_ = os.Mkdir("./pdf", os.ModePerm)
	}
	filename = "./pdf/" + filename
	//存储pdf文件
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		return
	}
	return
}

//解析文档
func parseUrl(url string) (content []byte, header http.Header, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	header = resp.Header
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return
}

func main() {
	var doi, temp string
	fmt.Println("请输入doi下载文献...或按回车键退出...")
	fmt.Printf("doi：")
	for {
		_, err := fmt.Scanln(&temp)
		if err != nil {
			break
		}
		if temp == doi {
			fmt.Println("重复的doi")
			continue
		}
		doi = temp
		filename, err := Download(doi)
		if err != nil {
			fmt.Println("下载错误", err)
		}
		filename = filepath.Base(filename)
		fmt.Println(filename, "下载完毕,请按回车键退出...")
	}
}
