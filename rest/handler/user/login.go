package user

import (
	"encoding/json"
	"net/http"

	"blog/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
	}

	accessToken, err := util.CreateJWT(h.cnf.JWTSecretKey, util.Payload{
		Sub:      usr.ID,
		FullName: usr.FullName,
		Email:    usr.Email,
		IsAuthor: usr.IsAuthor,
	})
	if err != nil {
		http.Error(w, "Internal server erro", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)
}
