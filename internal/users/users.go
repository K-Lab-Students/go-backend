package users

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/files"
	"lrm-backend/internal/models"
)

type UseCase struct {
	db   *sqlx.DB
	file *files.UseCase
}

func NewUsersUseCase(db *sqlx.DB, file *files.UseCase) *UseCase {
	return &UseCase{db: db, file: file}
}

func (uc *UseCase) GetUsers(limit, offset int) ([]models.User, error) {
	Users := make([]models.User, 0)
	if err := uc.db.Select(&Users, `select * from t_user limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range Users {
		if Users[i].FileObjectID != "" {
			filesPathes, err := uc.file.GetFiles(Users[i].FileObjectID)
			if err != nil {
				return nil, err
			}
			Users[i].Files = filesPathes
		}
	}

	return Users, nil
}
