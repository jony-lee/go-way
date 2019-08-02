package sci_hub

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func AnalysisDoi(key string) (doi string, err error) {
	a := strings.Split(doi, ".")[0]
	b := len(strings.Split(doi, "/"))
	//判断输入，若是doi则直接返回，若是字符串则进一步检索doi
	if a == "10" && b == 2 {
		return
	}
	//通过百度学术抓取检索第一篇文章的paperId
	BaseURL := "http://xueshu.baidu.com"
	url := BaseURL + "/s?wd=" + key
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	html := string(body)
	reg := regexp.MustCompile(`"\w{32}"`)
	paperId := reg.FindAllString(html, -1)
	if  paperId == nil {
		err = errors.New("未抓取到paperId")
		return
	}
	//通过paperId获取对应文章页面并抓取其中的doi
	paperUrl := fmt.Sprintf("http://xueshu.baidu.com/usercenter/paper/show?paperid=%s&site=xueshu_se", paperId[0])
	resp, err = http.Get(paperUrl)
	if err != nil {
		return
	}
	//doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the band and title
	//	band := s.Find("a").Text()
	//	title := s.Find("i").Text()
	if doiList := reg.FindAllString(html, -1); doiList != nil {
		doi = doiList[0]
		return
	}
	return
}
