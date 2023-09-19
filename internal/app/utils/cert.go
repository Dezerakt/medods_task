package utils

import (
	"encoding/base64"
	"os"
)

func EncodeToBase64(certContentSlice []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(certContentSlice)

	return encodeString
}

func GetWD() string {
	path, _ := os.Getwd()

	return path
}
