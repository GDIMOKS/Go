package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	GetRequest()
}



func GetRequest() {
	resp, err := http.Get("https://api.gu.spb.ru/mobile-adapter/doctor/GetDistrictList")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(string(body))

}