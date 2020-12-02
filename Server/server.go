package Server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandServer() {
	http.HandleFunc("/", handler)

	log.Println("Serving on http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request connection: %s, path: %s", r.Proto, r.URL.Path[1:])

	//fmt.Fprintf(w, "%s", r.Body)
	defer r.Body.Close()

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
