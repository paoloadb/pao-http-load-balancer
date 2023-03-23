package main

import (
	"os"
	"bufio"
	"log"
)

func GetIpList(fName string) []string {
	var list []string 

	file, err := os.Open(fName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		list = append(list, scan.Text())
	}

	return list
}