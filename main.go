package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)

var (
servers []string
activeServers []string
index = 0
globalCounter = 0
)

func main() {
	servers = GetIpList("sample-list.txt")
	fmt.Println(len(servers), " servers in list.")
	if len(servers) == 0 {
		log.Fatal("There's nothing on the server list")
	}
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func forwardRequest(res http.ResponseWriter, req *http.Request) {	
	activeServers = DoHealthCheck(servers)
	if len(activeServers) == 0 {
		fmt.Println(("No active servers!"))
		return
	}
	fmt.Println("globalcounter is at ", globalCounter)
	if globalCounter >= len(activeServers) {
		globalCounter = 0
	}

		u, err := url.Parse(activeServers[globalCounter])
		if err != nil {
			log.Fatal(err)
		}
		target := httputil.NewSingleHostReverseProxy(u)
		target.ServeHTTP(res, req)
	globalCounter +=1
}
