package sci_hub

import (
	"fmt"
	"net/http"
	"strings"
)

func analysisDoi(key string) (doi string, err error) {
	a := strings.Split(doi, ".")[0]
	b := len(strings.Split(doi, "/"))
	if a == "10" && b == 2 {
		return
	}
	BaseURL := "http://xueshu.baidu.com"
	url := BaseURL + "/s?wd=" + key
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
	return
}
