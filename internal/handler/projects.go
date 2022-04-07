package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetProjects(c *gin.Context) {
	//_, err := h.AuthUser(c)
	//if err != nil {
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

	Projects, err := h.Projects.GetProjects(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Projects)
}
func (h *Handler) GetProjectByID(c *gin.Context) {
	//_, err := h.AuthUser(c)
	//if err != nil {
	//	return
	//}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	ProjectModel, err := h.Projects.GetProject(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, ProjectModel)
}
func (h *Handler) AddProjects(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}

	var oneProjects models.Project
	if err := c.BindJSON(&oneProjects); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	oneProjects.AuthorID = user.ID
	oneProjects.FileObjectID = "{1}"

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

	ProjectModel, err := h.Projects.AddProject(&oneProjects)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, ProjectModel)
}
func (h *Handler) DeleteProjects(c *gin.Context) {
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

	if err := h.Projects.DeleteProject(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateProjects(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}

	var oneProjects models.Project
	if err := c.BindJSON(&oneProjects); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	oneProjects.ID = id
	oneProjects.AuthorID = user.ID
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

	ProjectModel, err := h.Projects.AddProject(&oneProjects)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, ProjectModel)
}
