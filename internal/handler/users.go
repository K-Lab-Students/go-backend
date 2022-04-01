package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetUsers(c *gin.Context) {
	var limit int
	var offset int

	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = 100
	}

	offset, err = strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	news, err := h.Users.GetUsers(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, news)
}
