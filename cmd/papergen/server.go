package main

import (
	"log"
	"papergen/internal/router"
)

func main() {
	r := router.Router()

	if err := r.Run(":1020"); err != nil {
		log.Fatal("server start error with msg: ", err.Error())
		return
	}
}
