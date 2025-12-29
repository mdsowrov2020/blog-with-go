package post

import (
	"fmt"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single post")
}
