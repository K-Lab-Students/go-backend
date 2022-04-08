package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"lrm-backend/internal/models"
	"lrm-backend/pkg/respfmt"
	"strconv"
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
	fmt.Println(token)
	fmt.Println(token2)
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

func getMd5ByUserStruct(md models.Md5) string {
	strmd5 := md.Email + md.Password + strconv.FormatBool(md.IsAdmin) + strconv.FormatBool(md.IsMember)
	fmt.Println("strmd5", strmd5)
	return strmd5
}

func (h *Handler) Registration(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		respfmt.BadRequest(c, "Данные пользователя не корректны")
		return
	}

	var a models.Auth
	if err := json.Unmarshal(jsonData, &a); err != nil {
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
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		respfmt.BadRequest(c, "Данные пользователя не корректны")
		return
	}

	var a models.Auth
	if err := json.Unmarshal(jsonData, &a); err != nil {
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
