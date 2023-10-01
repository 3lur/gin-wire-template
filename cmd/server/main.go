package main

import (
	"gin-mini-mall/pkg/http"
	"log"
)

func main() {
	logger := log.Logger{}
	app, cleanup, err := initApp(&logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, ":3000")

}
