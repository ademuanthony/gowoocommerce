package gowoocommerce

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

func GetUnixTime(micro bool) string {
	unixTime := time.Now().Unix()
	unixTimeStr := strconv.FormatInt(unixTime, 10)

	if !micro {
		return unixTimeStr[:10]
	}

	return "0." + unixTimeStr[10:16] + "00" + unixTimeStr[:10]
}

func getSha256(key, message string) string {
	secret := []byte(key)
	messageByte := []byte(message)

	hash := hmac.New(sha256.New, secret)
	hash.Write(messageByte)
	//return hex.EncodeToString(hash.Sum(nil))

	hashString := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return hashString
}
