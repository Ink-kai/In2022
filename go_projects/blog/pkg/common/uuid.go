package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Uuid_generate_v3() string {
	slat := []byte(time.Now().UTC().String())
	h := md5.New()
	h.Write(slat)
	return hex.EncodeToString(h.Sum(nil))
}
func Uuid_generate_password(password string) string {
	slat := []byte(time.Now().UTC().String() + password)
	h := md5.New()
	h.Write(slat)
	return hex.EncodeToString(h.Sum(nil))
}
