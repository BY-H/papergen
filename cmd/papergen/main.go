package main

import (
	"log"
	"papergen/internal/global"
	_ "papergen/internal/global"
	"papergen/internal/router"
)

func main() {
	r := router.Router()

	if err := r.Run(global.Conf.Port); err != nil {
		log.Fatal("server start error with msg: ", err.Error())
		return
	}
}
