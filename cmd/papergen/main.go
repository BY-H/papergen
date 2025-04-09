package main

import (
	"log"
	"papergen/internal/global"
	_ "papergen/internal/global"
	"papergen/internal/router"
	"strconv"
)

func main() {
	r := router.Router()

	if err := r.Run(strconv.Itoa(global.Conf.Port)); err != nil {
		log.Fatal("server start error with msg: ", err.Error())
		return
	}
}
