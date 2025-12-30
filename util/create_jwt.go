package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"jwt"`
}

type Payload struct {
	Sub      int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsAuthor bool   `json:"is_author"`
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

	base64Header := Base64UrlEncoder(byteHeaderArr)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	base64Data := Base64UrlEncoder(byteArrData)

	message := base64Header + "." + base64Data

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)

	base64Signature := Base64UrlEncoder(signature)

	jwt := base64Header + "." + base64Data + "." + base64Signature

	return jwt, nil
}
