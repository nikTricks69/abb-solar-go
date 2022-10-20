package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/jsonquery"
)

// curl http://user21:password@192.168.80.172/v1/livedata
func main() {
	log.Print("Starting..")
	http.HandleFunc("/metrics", getMetrics())
	log.Fatal(http.ListenAndServe(fmt.Sprint(":8869"), nil))
}
func getMetrics() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			userName := os.Getenv("abb_user")
			password := os.Getenv("abb_password")
			url := os.Getenv("abb_ip")
			//log.Print(url)
			endpoint := strings.ReplaceAll(fmt.Sprintf("http://%s:%s@ %s/v1/livedata", userName, password, url), " ", "")
			doc, err := jsonquery.LoadURL(endpoint)
			if err != nil {
				log.Fatal("Endpoint failed %s", endpoint)

				panic(err)
			}

			list := jsonquery.Find(doc, "//points/*")
			for _, n := range list {

				myMap := n.Value().(map[string]interface{})
				key := fmt.Sprint(myMap["name"])
				val := fmt.Sprint(myMap["value"])
				w.Write([]byte(key))
				w.Write([]byte(" "))
				w.Write([]byte(val))
				w.Write([]byte("\n"))
				//YAHOOO
				//fmt.Println(n.Value().name)
			}
		}

	})
}
