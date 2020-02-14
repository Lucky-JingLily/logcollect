package main

import "fmt"

func main() {
	filename := "conf/logagent.conf"
	err := loadConf(filename)
	if err != nil {
		fmt.Println("load conf failed, err:", err)
		panic("load conf failed")
		return
	}
}
