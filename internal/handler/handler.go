package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lrm-backend/internal/competitions"
	"lrm-backend/internal/faculties"
	"lrm-backend/internal/news"
	"lrm-backend/internal/projects"
	studyplace "lrm-backend/internal/study_place"
	"lrm-backend/internal/users"
)

type Handler struct {
	News         *news.UseCase
	Projects     *projects.UseCase
	Users        *users.UseCase
	Competitions *competitions.UseCase
	Faculties    *faculties.UseCase
	StudyPlace   *studyplace.UseCase
}

func New(News *news.UseCase, Projects *projects.UseCase, Users *users.UseCase, Competitions *competitions.UseCase, Faculties *faculties.UseCase, StudyPlace *studyplace.UseCase) *Handler {
	return &Handler{
		News:         News,
		Projects:     Projects,
		Users:        Users,
		Competitions: Competitions,
		Faculties:    Faculties,
		StudyPlace:   StudyPlace,
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

	r.GET("competitions", h.GetCompetitions)
	r.GET("competitions/:id", h.GetCompetitionByID)
	r.POST("competitions", h.AddCompetitions)
	r.DELETE("competitions/:id", h.DeleteCompetitions)
	r.PUT("competitions/:id", h.UpdateCompetitions)

	r.GET("faculties", h.GetFaculties)
	r.GET("faculties/:id", h.GetFacultyByID)
	r.POST("faculties", h.AddFaculties)
	r.DELETE("faculties/:id", h.DeleteFaculties)
	r.PUT("faculties/:id", h.UpdateFaculties)

	r.GET("studyPlaces", h.GetStudyPlaces)
	r.GET("studyPlaces/:id", h.GetStudyPlaceByID)
	r.POST("studyPlaces", h.AddStudyPlaces)
	r.DELETE("studyPlaces/:id", h.DeleteStudyPlace)
	r.PUT("studyPlaces/:id", h.UpdateStudyPlace)

	r.GET("users", h.GetUsers)
	r.GET("user", h.GetUserByID)
	r.PUT("user/update", h.UpdateUser)

	r.POST("register", h.Registration)
	r.POST("login", h.Login)

	return r
}
