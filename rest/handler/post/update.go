package post

import (
	"fmt"
	"net/http"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update post")
}
