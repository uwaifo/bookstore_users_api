package utils

import (
	"crypto"
	//"crypto/md5"
	"encoding/hex"
)

//GetMd5 . . .
func GetMd5(input string) string {
	hash := crypto.MD5.New()
	// md5.New()

	//hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}
