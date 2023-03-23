package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)

var (
servers = []string {
	"http://localhost:9001",
	"http://localhost:9000",
	"http://localhost:9002",
}
index = 0
globalCounter = 0
)

func main() {

		fmt.Println(len(servers), " servers.")
		http.HandleFunc("/", forwardRequest)
		log.Fatal(http.ListenAndServe(":8080", nil))

}

func forwardRequest(res http.ResponseWriter, req *http.Request) {	
	fmt.Println("globalcounter is at ", globalCounter)
	if globalCounter == 3 {
		globalCounter = 0
	}
		u := getServerFromList()
		target := httputil.NewSingleHostReverseProxy(u)
		target.ServeHTTP(res, req)
	globalCounter +=1
}

func getServerFromList() *url.URL {
	ctr := index % len(servers)
	fmt.Println("counter at ", ctr)
	x, err := url.Parse(servers[ctr])
	if err != nil {
		panic(err)
	}
	index++
	fmt.Println("Serving", servers[ctr])
	return x
}