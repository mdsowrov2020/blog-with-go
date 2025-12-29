package post

import (
	"fmt"
	"net/http"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all post")
}
