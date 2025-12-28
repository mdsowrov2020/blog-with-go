package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(statusCode)
}

func SendError(w http.ResponseWriter, statusCode int, data string) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(statusCode)
}
