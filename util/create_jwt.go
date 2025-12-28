package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"jwt"`
}

type Payload struct {
	Sub       int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteHeaderArr, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
	}

	base64Header := base64Convert(byteHeaderArr)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	base64Data := base64Convert(byteArrData)

	message := base64Header + "." + base64Data

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)

	base64Signature := base64Convert(signature)

	jwt := base64Header + "." + base64Data + "." + base64Signature

	return jwt, nil
}

func base64Convert(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
