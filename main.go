package main

import (
	"net/http"

	"github.com/hebelsan/gohttp/pkg/middleware"
	"github.com/hebelsan/gohttp/pkg/util"
)

func main() {
	http.Handle("/", middleware.FileHandler{})
	err := http.ListenAndServe(util.GetPort("80"), nil)
	if err != nil {
		panic(err)
	}
}
