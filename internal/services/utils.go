package services

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	"net/url"
)

func randomStr(length int) string {
	buff := make([]byte, int(math.Ceil(float64(length)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:length]
}

func UrlIsValid(text string) error {
	_, err := url.ParseRequestURI(text)
	return err
}
