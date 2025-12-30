package post

import (
	"net/http"

	"blog/util"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.postRepo.List()
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, list)
}
