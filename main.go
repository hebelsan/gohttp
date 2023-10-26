// Package main starts the server to serve static files and provide upload functionality
package main

import (
	"net/http"
	"os"

	"github.com/hebelsan/gohttp/pkg/auth"
	"github.com/hebelsan/gohttp/pkg/file"
	"github.com/hebelsan/gohttp/pkg/util"
)

func main() {
	authMiddleware := auth.NewMiddleware(os.Getenv(auth.ENV_KEY))
	http.HandleFunc("/", authMiddleware.AuthHandler(file.Handler))
	err := http.ListenAndServe(util.GetPort("80"), nil)
	if err != nil {
		panic(err)
	}
}
