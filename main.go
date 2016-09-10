package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	uri := os.Getenv("BB_PROXY_URI")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("err:", err)
			return
		}

		query, err := url.ParseQuery(string(body))
		if err != nil {
			log.Println("err:", err)
			return
		}

		urlStr := uri + "?" + query.Encode()
		log.Println(urlStr)

		_, err = http.Get(urlStr)
		if err != nil {
			log.Println(err)
			return
		}
	})

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
