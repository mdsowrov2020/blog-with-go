package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blog/domain"
	"blog/util"
)

type UpdatePostReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	pID, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid post id")
		return
	}

	var req UpdatePostReq

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid request body")
		return
	}

	_, err = h.svc.Update(domain.Post{
		ID:          pID,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully updated post")
}
