package handler

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) AuthUser(c *gin.Context) (models.User, error) {
	token := c.GetHeader("x-token")
	userIDstr := c.GetHeader("user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		respfmt.BadRequest(c, "user_id is not correct")
		return models.User{}, err
	}
	user, err := h.Users.GetUserByID(userID)
	if err != nil {
		respfmt.BadRequest(c, "user_id is not correct")
		return models.User{}, err
	}
	token2 := getMd5ByUserStruct(models.Md5{
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
		IsMember: user.IsMember,
	})
	if err := bcrypt.CompareHashAndPassword([]byte(token), []byte(token2)); err != nil {
		respfmt.BadRequest(c, "you token is not correct")
		return models.User{}, errors.New("you token is not correct")
	}

	userRole := 2
	if user.IsAdmin {
		userRole = 0
	} else if user.IsMember {
		userRole = 1
	}
	user.UserRoleID = userRole
	return user, nil
}

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
func (h *Handler) GetUserByID(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}

	news, err := h.Users.GetUserByID(user.ID)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, news)
}
func (h *Handler) UpdateUser(c *gin.Context) {
	user, err := h.AuthUser(c)
	if err != nil {
		return
	}
	var profile models.UserUpdateProfile
	if err := c.BindJSON(&profile); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	profile.ID = user.ID

	fileName := strconv.Itoa(int(time.Now().Unix()) + rand.Intn(100000))
	if profile.File != nil && *profile.File != "" {
		fmt.Println(1)
		filePath, err := saveImageToDisk(fileName, *profile.File)
		if err != nil {
			respfmt.BadRequest(c, err.Error())
			return
		}
		fmt.Println(2)
		id, err := h.Users.SaveUserFile(filePath)
		if err != nil {
			respfmt.BadRequest(c, err.Error())
			return
		}
		fmt.Println("id:", id)
		profile.FileID = &id
	}

	if profile.FileID == nil {
		profile.FileID = user.FileObjectID
	}

	if profile.CompetitionsID == nil {
		profile.CompetitionsID = user.CompetitionsID
	}

	fmt.Printf("\n%+v\n", profile)

	if err := h.Users.UpdateUserProfile(&profile); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, "ok")
}
func saveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", fmt.Errorf("is not correct file")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	fileName := "./photo/" + fileNameBase + ".jpeg"
	if err := ioutil.WriteFile(fileName, buff.Bytes(), 0777); err != nil {
		return "", err
	}

	return "/photo/" + fileNameBase + ".jpeg", err
}
func getMd5ByUserStruct(md models.Md5) string {
	strmd5 := md.Email + md.Password + strconv.FormatBool(md.IsAdmin) + strconv.FormatBool(md.IsMember)
	return strmd5
}

func (h *Handler) Registration(c *gin.Context) {
	var a models.Auth
	if err := c.BindJSON(&a); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	user, err := h.Users.Registration(&a)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	token := getMd5ByUserStruct(models.Md5{
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
		IsMember: user.IsMember,
	})

	fmt.Println("t1", token)
	userRole := 2
	if user.IsAdmin {
		userRole = 0
	} else if user.IsMember {
		userRole = 1
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	fmt.Println("t2", string(hashedPassword))

	respfmt.OK(c, map[string]interface{}{
		"token":     string(hashedPassword),
		"user_role": userRole,
		"id":        user.ID,
	})
}
func (h *Handler) Login(c *gin.Context) {
	var a models.Auth
	if err := c.BindJSON(&a); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	fmt.Printf("\n%+v\n", a)

	user, err := h.Users.Login(&a)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	token := getMd5ByUserStruct(models.Md5{
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
		IsMember: user.IsMember,
	})

	userRole := 2
	if user.IsAdmin {
		userRole = 0
	} else if user.IsMember {
		userRole = 1
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"token":     string(hashedPassword),
		"user_role": userRole,
		"id":        user.ID,
	})
}
