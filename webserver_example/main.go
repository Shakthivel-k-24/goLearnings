package main

import (
	"fmt"
	"sample-server/main/envfetcher"
	"strconv"
)

func main() {

	var data envfetcher.Data
	strvar := "1024"
	val, err := strconv.ParseInt(strvar, 10, 64)
	defer fmt.Println(fmt.Sprintf("%d,%v", val, err))
	fmt.Println("Hello there")
	panic(data)
}
