package Server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ClientPost() {
	client := http.Client{}

	resp, err := client.Post("http://127.0.0.1:8080/hello", "text/plain", bytes.NewBufferString("\nHello Server!"))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))
}
