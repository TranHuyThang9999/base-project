package main

import (
	"flag"
	"log"
	"rices/common/configs"
	"rices/common/utils"
)

func init() {
	var pathConfig string
	flag.StringVar(&pathConfig, "configs", "configs/configs.json", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig)
}

func main() {

	err := utils.SendEmail("tranhuythang9999@gmail.com", "test", "hio")
	if err != nil {
		log.Println(err)
	}
	log.Println("send")
}
