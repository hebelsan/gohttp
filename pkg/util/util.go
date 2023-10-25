// Package util implements utility routines such as environment variables handling
package util

import (
	"crypto/rand"
	"fmt"
	"os"
)

// GetPort retrieves the port from env or returns the default port
func GetPort(defaultPort string) string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return ":" + defaultPort
}

// PseudoUuid creates and returns a pseudo UUID by utilizing crypto/rand
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
