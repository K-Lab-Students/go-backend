package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetStudyPlaces(c *gin.Context) {
	//if _, err := h.AuthUser(c); err != nil {
	//	return
	//}

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

	Facultys, err := h.StudyPlace.GetStudyPlaces(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Facultys)
}
func (h *Handler) GetStudyPlaceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	FacultyModel, err := h.StudyPlace.GetStudyPlace(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, FacultyModel)
}
func (h *Handler) AddStudyPlaces(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneStudyPlace models.StudyPlace
	if err := c.BindJSON(&oneStudyPlace); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	FacultyModel, err := h.StudyPlace.AddStudyPlace(&oneStudyPlace)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, FacultyModel)
}
func (h *Handler) DeleteStudyPlace(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	if err := h.StudyPlace.DeleteStudyPlace(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateStudyPlace(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneStudyPlace models.StudyPlace
	if err := c.BindJSON(&oneStudyPlace); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	oneStudyPlace.ID = id

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

	FacultyModel, err := h.StudyPlace.AddStudyPlace(&oneStudyPlace)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, FacultyModel)
}
