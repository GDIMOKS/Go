package Server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest() {
	resp, err := http.Get("http://127.0.0.1:8080/")
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(body))
}
