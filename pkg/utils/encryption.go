package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string, salt ...string) string {
	h := md5.New()
	h.Write([]byte(s))
	if len(salt) > 0 {
		h.Write([]byte(salt[0]))
	}
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}
