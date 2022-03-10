package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(value string) string {
	data := []byte(value)
	md5Cxt := md5.New()
	md5Cxt.Write(data)
	return hex.EncodeToString(md5Cxt.Sum(nil))
}

func MD5Salt(value string, salt string) string {
	return MD5(value + salt)
}

func MD5Count(value string, count int) string {
	for i := 0; i < count; i++ {
		value = MD5(value)
	}
	return value
}

func MD5SaltCount(value string, salt string, count int) string {
	for i := 0; i < count; i++ {
		if i == 0 {
			value = MD5Salt(value, salt)
		} else {
			value = MD5(value)
		}
	}
	return value
}
