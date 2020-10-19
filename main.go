package main

import (
	"be04gomy/config"
	"fmt"
)

func main() {
	_, err := config.ConnectMysql()
	
	fmt.Println(err)
}
