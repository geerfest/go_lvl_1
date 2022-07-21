package main

import (
	"read_conf/conf_reade_json"
)

func main() {
	//config := conf_reader.GetConfig()
	//config.Print()
	conf := conf_reade_json.GetConfig("conf.json")
	conf.Print()
}
