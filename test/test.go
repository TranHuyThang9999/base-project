package main

import (
	"log"
	"rices/common/utils"
)

func main() {
	err := utils.SendEmail("tranhuythang9999@gmail.com", "test", "hio")
	log.Println(err)
}
