package user

import (
	"encoding/json"
	"net/http"

	"blog/domain"
	"blog/util"
)

type ReqUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAuthor bool   `json:"is_author"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req ReqUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid request body")
		return
	}

	user, err := h.svc.Create(domain.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		IsAuthor: req.IsAuthor,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, user)
}
