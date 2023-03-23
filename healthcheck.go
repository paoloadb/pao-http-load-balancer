package main

import (
	"net"
	"time"
	// "fmt"
	"strings"
)

func DoHealthCheck(addr []string) []string{
	var newList []string
	for _, j := range addr {
		frmttdstr := strings.Replace(j, "http://", "", 1)
		// fmt.Println("Dialing ", frmttdstr)
 		_, err := net.DialTimeout("tcp", frmttdstr, 1 * time.Millisecond)	
		if err == nil {
			newList = append(newList, j)
		}
	}
	return newList
}