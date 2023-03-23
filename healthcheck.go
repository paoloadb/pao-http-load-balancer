package main

import (
	"net"
	"time"
)

func DoHealthCheck(addr string) bool {
	_, err := net.DialTimeout("tcp", addr, 5 * time.Millisecond)
	if err != nil {
		return false 
	}
	return true
}