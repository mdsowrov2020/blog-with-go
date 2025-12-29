package post

import (
	"fmt"
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete post")
}
