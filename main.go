package main

import (
	"net/http"

	"gohttp/pkg/middleware"
	"gohttp/pkg/util"
)

func main() {
	http.Handle("/", middleware.FileHandler{})
	err := http.ListenAndServe(util.GetPort("80"), nil)
	if err != nil {
		panic(err)
	}
}
