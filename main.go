package main

import (
	"fmt"
	"github.com/jony-lee/go-way/sci-hub"
)

func main() {
	doi, err := sci_hub.AnalysisDoi("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doi)
}
