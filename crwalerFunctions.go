package main

import (
	"fmt"
	"io"
	"net/http"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func crwalForsythCounty(address string) {
	var url string = "https://lrcpwa.ncptscloud.com/forsyth/searchProperty.aspx"
	getInitials(url)
}
func getInitials(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erorr %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	content := string(body)
	println(content)

}
