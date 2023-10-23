package util

import "os"

func GetPort(defaultPort string) string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return ":" + defaultPort
}
