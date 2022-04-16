package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const photoLink = "http://localhost:8080"

func (h *Handler) GetNews(c *gin.Context) {
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

	news, err := h.News.GetNews(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, news)
}
func (h *Handler) GetNewByID(c *gin.Context) {
	//if _, err := h.AuthUser(c); err != nil {
	//	return
	//}

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
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneNews models.New
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	//if err := c.BindJSON(&oneNews); err != nil {
	if err := json.Unmarshal(b, &oneNews); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if oneNews.ID != 0 {
		if err := h.News.DeleteFiles(oneNews.ID, []string{}); err != nil {
			respfmt.InternalServer(c, err.Error())
		}
	}

	var fileObjects []string
	for {
		i := strings.Index(oneNews.Body, "data:")
		var b []byte
		var j int
		if i != -1 {
			j = i
			for string(oneNews.Body[j]) != "\"" {
				b = append(b, oneNews.Body[j])
				j++
			}
		} else {
			break
		}

		fileName := strconv.Itoa(int(time.Now().Unix()) + rand.Intn(100000000))
		filePath, err := saveImageToDisk(fileName, string(b))
		if err != nil {
			respfmt.BadRequest(c, err.Error())
			return
		}
		id, err := h.Users.SaveUserFile(filePath)
		if err != nil {
			respfmt.BadRequest(c, err.Error())
			return
		}

		fileObjects = append(fileObjects, strconv.Itoa(id))

		q := append([]byte(oneNews.Body)[:i], []byte(photoLink+filePath)...)
		q = append(q, []byte(oneNews.Body)[j:]...)

		oneNews.Body = string(q)
	}
	oneNews.AuthorID = user.ID
	oneNews.FileObjectID = "{" + strings.Join(fileObjects, ",") + "}"

	newModel, err := h.News.AddNew(&oneNews)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, newModel)
}
func (h *Handler) DeleteNews(c *gin.Context) {
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

	if err := h.News.DeleteFiles(id, []string{}); err != nil {
		respfmt.InternalServer(c, err.Error())
	}
	if err := h.News.DeleteNew(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func (h *Handler) UpdateNews(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	if user.UserRoleID == 2 {
		respfmt.BadRequest(c, "user have not permission")
		return
	}
	var oneNews models.New
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	//if err := c.BindJSON(&oneNews); err != nil {
	if err := json.Unmarshal(b, &oneNews); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "not correct id")
		return
	}

	oneNews.ID = id

	var fileObjects []string
	var fileID int
	var dontDeleteFiles []string
	qw := oneNews.Body

	for {
		i := strings.Index(oneNews.Body, "data:")
		var b []byte
		var j int
		if i != -1 {
			j = i
			for string(oneNews.Body[j]) != "\"" {
				b = append(b, oneNews.Body[j])
				j++
			}
			fileName := strconv.Itoa(int(time.Now().Unix()) + rand.Intn(100000000))
			filePath, err := saveImageToDisk(fileName, string(b))
			if err != nil {
				respfmt.BadRequest(c, err.Error())
				return
			}
			fileID, err = h.Users.SaveUserFile(filePath)
			if err != nil {
				respfmt.BadRequest(c, err.Error())
				return
			}
			q := append([]byte(oneNews.Body)[:i], []byte(photoLink+filePath)...)
			q = append(q, []byte(oneNews.Body)[j:]...)

			oneNews.Body = string(q)

		} else {
			i := strings.Index(qw, "/photo")
			b = []byte{}
			j = 0
			if i != -1 {
				j = i
				for string(qw[j]) != "\"" {
					b = append(b, qw[j])
					j++
				}
			} else {
				break
			}

			dontDeleteFiles = append(dontDeleteFiles, string(b))

			q := append([]byte(qw)[:i], []byte(qw)[j:]...)
			qw = string(q)
			fileID, err = h.Users.SearchFile(string(b))
			if err != nil {
				respfmt.InternalServer(c, err.Error())
			}
		}

		fileObjects = append(fileObjects, strconv.Itoa(fileID))

	}

	if oneNews.ID != 0 {
		if err := h.News.DeleteFiles(oneNews.ID, dontDeleteFiles); err != nil {
			respfmt.InternalServer(c, err.Error())
		}
	}

	oneNews.FileObjectID = "{" + strings.Join(fileObjects, ",") + "}"
	oneNews.AuthorID = user.ID

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
