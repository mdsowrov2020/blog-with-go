package post

import (
	"net/http"
	"strconv"

	"blog/util"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	pID, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid post id")
		return
	}

	blog, err := h.postRepo.Get(pID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server error")
		return
	}

	if blog == nil {
		util.SendError(w, http.StatusNotFound, "Post not found")
		return
	}

	util.SendData(w, http.StatusOK, blog)
}
