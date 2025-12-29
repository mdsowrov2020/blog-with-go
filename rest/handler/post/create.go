package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"blog/util"
)

type PostReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"img_url"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create POST")

	var postReq PostReq

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&postReq)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid request body")
		return
	}
}
