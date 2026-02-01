package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPasswordMD5(password string) (string, error) {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:]), nil
}

func HashPasswordBase64(password string) (string, error) {
	hash := []byte(password)

	return hex.EncodeToString(hash), nil
}
