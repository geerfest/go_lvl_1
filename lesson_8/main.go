package main

import (
	"read_conf/conf_reader"
)

func main() {
	config := conf_reader.GetConfig()
	config.Print()
}
