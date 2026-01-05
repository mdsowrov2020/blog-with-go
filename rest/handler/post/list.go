package post

import (
	"net/http"
	"strconv"

	"blog/util"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	reqPageStr := reqQuery.Get("page")
	reqLimitStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(reqPageStr, 10, 32)
	limit, _ := strconv.ParseInt(reqLimitStr, 10, 32)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}
	list, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	util.SendPage(w, list, page, limit, cnt)
}
