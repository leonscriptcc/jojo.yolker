package main

import (
	"github.com/leonscriptcc/jojo.yolker/config"
	"github.com/leonscriptcc/jojo.yolker/service"
	"log"
)

func init() {
	if err := config.Load(); err != nil {
		log.Panic("load config fail:", err)
	}
}

func main() {

	if err := service.WriteV(); err != nil {
		log.Println("writev fail:", err.Error())
		return
	}

	log.Println("Hello! my honey~")
}
