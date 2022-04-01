package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/news"
	"lrm-backend/internal/projects"
	"lrm-backend/internal/users"
)

type Handler struct {
	News     *news.UseCase
	Projects *projects.UseCase
	Users    *users.UseCase
}

func New(News *news.UseCase, Projects *projects.UseCase, Users *users.UseCase) *Handler {
	return &Handler{
		News:     News,
		Projects: Projects,
		Users:    Users,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Static("/photo", "./photo")

	r.Use(cors.Default())
	r.Use(AuthUser)

	r.GET("news", h.GetNews)
	r.GET("news/:id", h.GetNewByID)
	r.POST("news", h.AddNews)
	r.DELETE("news/:id", h.DeleteNews)
	r.PUT("news/:id", h.UpdateNews)

	r.GET("projects", h.GetProjects)
	r.GET("projects/:id", h.GetProjectByID)
	r.POST("projects", h.AddProjects)
	r.DELETE("projects/:id", h.DeleteProjects)
	r.PUT("projects/:id", h.UpdateProjects)

	r.GET("users", h.GetUsers)
	r.GET("users/:id")
	r.DELETE("users/:id")
	r.PUT("users/:id")

	r.POST("register")
	r.POST("login")
	return r
}

func AuthUser(c *gin.Context) {
	c.Set("user_id", 1)
}

func GetAuthUser(c *gin.Context) int {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	userIdInt, exists := userID.(int)
	if !exists {
		return 0
	}

	return userIdInt
}
