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
	//r.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "user_id", "x-token")
	r.Use(cors.New(config))
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

	r.POST("register", h.Registration)
	r.POST("login", h.Login)

	return r
}
