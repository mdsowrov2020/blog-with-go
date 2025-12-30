package post

import (
	"net/http"
	"strconv"

	"blog/util"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	pID, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid blog id")
		return
	}

	err = h.postRepo.Delete(pID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal server error")
		return
	}

	util.SendData(w, "Successfully deleted product", http.StatusOK)
}
