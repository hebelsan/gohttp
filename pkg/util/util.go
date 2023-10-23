package util

import (
	"crypto/rand"
	"fmt"
	"os"
)

func GetPort(defaultPort string) string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return ":" + defaultPort
}

// Note - NOT RFC4122 compliant
func PseudoUuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
