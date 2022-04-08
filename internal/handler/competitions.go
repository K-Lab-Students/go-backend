package handler

import (
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
)

func (h *Handler) GetCompetitions(c *gin.Context) {
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

	Competitions, err := h.Competitions.GetCompetitions(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Competitions)
}
func (h *Handler) GetCompetitionByID(c *gin.Context) {
	//if _, err := h.AuthUser(c); err != nil {
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		respfmt.BadRequest(c, "not correct id")
	}

	CompetitionModel, err := h.Competitions.GetCompetition(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, CompetitionModel)
}
func (h *Handler) AddCompetitions(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneCompetitions models.Competition
	if err := c.BindJSON(&oneCompetitions); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	oneCompetitions.FileObjectID = "{1}"
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

	CompetitionModel, err := h.Competitions.AddCompetition(&oneCompetitions)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, CompetitionModel)
}
func (h *Handler) DeleteCompetitions(c *gin.Context) {
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

	if err := h.Competitions.DeleteCompetition(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateCompetitions(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneCompetitions models.Competition
	if err := c.BindJSON(&oneCompetitions); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	oneCompetitions.ID = id

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

	CompetitionModel, err := h.Competitions.AddCompetition(&oneCompetitions)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, CompetitionModel)
}
