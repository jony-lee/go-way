package main

import (
	"fmt"
	"github.com/jony-lee/go-way/sci-hub/paper"
	"io"
	"net/http"
)

func run() {
	http.HandleFunc("/paperDownload", paperDownloadHandler)
	_ = http.ListenAndServe("localhost:5000", nil)
}

func paperDownloadHandler(w http.ResponseWriter, r *http.Request) {
	doi := r.PostFormValue("doi")
	mail := r.PostFormValue("mail")
	paperDownload(doi, mail)
	resp := fmt.Sprintf("已收到请求【doi: %s ,mail: %s】\n", doi, mail)
	_, _ = io.WriteString(w, resp)
	return
}

func paperDownload(doi string, mail string) {
	msg, err := paper.Download(doi, mail)
	if err != nil {

	}
}

func main() {
	run()
}
