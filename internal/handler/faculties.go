package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetFaculties(c *gin.Context) {
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

	Facultys, err := h.Faculties.GetFacultys(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Facultys)
}
func (h *Handler) GetFacultyByID(c *gin.Context) {
	//if _, err := h.AuthUser(c); err != nil {
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	FacultyModel, err := h.Faculties.GetFaculty(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, FacultyModel)
}
func (h *Handler) AddFaculties(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneFacultys models.Faculty
	if err := c.BindJSON(&oneFacultys); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
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

	FacultyModel, err := h.Faculties.AddFaculty(&oneFacultys)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, FacultyModel)
}
func (h *Handler) DeleteFaculties(c *gin.Context) {
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

	if err := h.Faculties.DeleteFaculty(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateFaculties(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneFacultys models.Faculty
	if err := c.BindJSON(&oneFacultys); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	oneFacultys.ID = id

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

	FacultyModel, err := h.Faculties.AddFaculty(&oneFacultys)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, FacultyModel)
}
