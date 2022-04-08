package main

import (
	"go.uber.org/zap"
	"lrm-backend/internal/competitions"
	Facultys "lrm-backend/internal/faculties"
	"lrm-backend/internal/files"
	"lrm-backend/internal/handler"
	"lrm-backend/internal/news"
	"lrm-backend/internal/projects"
	studyplace "lrm-backend/internal/study_place"
	"lrm-backend/internal/users"
	"lrm-backend/pkg/database"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db, err := database.Connect("postgres", "postgres", "postgres", "localhost", "5432")
	if err != nil {
		logger.Fatal("db connect", zap.Error(err))
	}

	fileUC := files.NewUseCase(db)

	studyPlaceUC := studyplace.NewStudyPlaceUseCase(db)
	facultiesUC := Facultys.NewFacultyUseCase(db)
	competitionsUC := competitions.NewCompetitionUseCase(db, fileUC)
	usersUC := users.NewUsersUseCase(db, fileUC, competitionsUC, facultiesUC)
	newsUC := news.NewUseCase(db, fileUC)
	projectsUC := projects.NewProjectUseCase(db, fileUC)

	h := handler.New(newsUC, projectsUC, usersUC, competitionsUC, facultiesUC, studyPlaceUC)
	server := h.InitRoutes()

	server.Run()
}
