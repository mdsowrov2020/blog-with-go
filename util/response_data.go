package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
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
