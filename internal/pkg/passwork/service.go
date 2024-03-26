package passwork

import (
	"crypto/md5"
	"fmt"
)

func GetHash(password string) string {
	hash := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", hash[:])
}

func CheckHash(password string, hash string) bool {
	return GetHash(password) == hash
}
