package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetNews(c *gin.Context) {
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

	news, err := h.News.GetNews(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, news)
}
func (h *Handler) GetNewByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	newModel, err := h.News.GetNew(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, newModel)
}
func (h *Handler) AddNews(c *gin.Context) {
	var oneNews models.New
	if err := c.BindJSON(&oneNews); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	userID := GetAuthUser(c)
	if userID == 0 {
		respfmt.BadRequest(c, "not auth")
	}
	oneNews.AuthorID = userID
	oneNews.FileObjectID = "{1}"
	//formFile, err := c.FormFile("file")
	//if err != nil {
	//	respfmt.BadRequest(c, err.Error())
	//	return
	//}
	//file, err := formFile.Open()
	//if err != nil {
	//	respfmt.BadRequest(c, err.Error())
	//	return
	//}

	newModel, err := h.News.AddNew(&oneNews)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, newModel)
}
func (h *Handler) DeleteNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	if err := h.News.DeleteNew(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateNews(c *gin.Context) {
	var oneNews models.New
	if err := c.BindJSON(&oneNews); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	userID := GetAuthUser(c)
	if userID == 0 {
		respfmt.BadRequest(c, "not auth")
		return
	}
	oneNews.ID = id
	oneNews.AuthorID = userID

	//formFile, err := c.FormFile("file")
	//if err != nil {
	//	respfmt.BadRequest(c, err.Error())
	//	return
	//}
	//file, err := formFile.Open()
	//if err != nil {
	//	respfmt.BadRequest(c, err.Error())
	//	return
	//}

	newModel, err := h.News.AddNew(&oneNews)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, newModel)
}
