package post

import (
	"encoding/json"
	"net/http"

	"blog/repo"
	"blog/util"
)

type PostReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var postReq PostReq

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&postReq)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid request body")
		return
	}

	createdPost, err := h.postRepo.Create(repo.Post{
		Title:       postReq.Title,
		Description: postReq.Description,
		ImageURL:    postReq.ImageURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, createdPost)
}
