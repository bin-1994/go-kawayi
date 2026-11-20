package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const AUTH_KEY = "U2FsdGVkX18Wv2TVLmvrnKZ3mGRb0XdlDxprEoKQEOU="

func Md5(value string) string {
	encrypt := md5.Sum([]byte(value))
	return fmt.Sprintf("%x", encrypt)
}

func SHA1(value string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(value))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

func UserMd5(value string) string {
	str := SHA1(value) + AUTH_KEY
	return Md5(str)
}

func CurrentTimestamp() int64 {
	year, month, day := time.Now().Date()
	timestamp := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	return timestamp
}

func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
