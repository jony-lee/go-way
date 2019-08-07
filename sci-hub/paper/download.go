package main

import (
	"fmt"
	"time"
)

func Download(doi, mail string) (msg string, err error) {
	time.Sleep(time.Second * 10)
	fmt.Println("下载完成", doi)
	return
}
