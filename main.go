package main

import (
	"net/http"
	"os"

	"github.com/hebelsan/gohttp/pkg/auth"
	"github.com/hebelsan/gohttp/pkg/file"
	"github.com/hebelsan/gohttp/pkg/util"
)

func main() {
	handler := file.Handler
	if os.Getenv("AUTH") == "API-KEY" {
		authMiddleware := auth.NewMiddleware()
		handler = authMiddleware.Handle(file.Handler)
	}
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(util.GetPort("80"), nil)
	if err != nil {
		panic(err)
	}
}
