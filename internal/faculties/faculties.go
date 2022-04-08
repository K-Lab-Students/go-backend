package faculties

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/models"
	"strings"
)

type UseCase struct {
	db *sqlx.DB
}

func NewFacultyUseCase(db *sqlx.DB) *UseCase {
	return &UseCase{db: db}
}

func (uc *UseCase) GetUserFaculties(faculties string) ([]models.Faculty, error) {
	faculties = strings.ReplaceAll(faculties, "{", "")
	faculties = strings.ReplaceAll(faculties, "}", "")

	userFacultys := make([]models.Faculty, 0)
	if err := uc.db.Select(&userFacultys, `select id, name, color from t_faculties where id in (`+faculties+`)`); err != nil {
		return nil, err
	}
	return userFacultys, nil
}

func (uc *UseCase) GetFacultys(limit, offset int) ([]models.Faculty, error) {
	Facultys := make([]models.Faculty, 0)
	if err := uc.db.Select(&Facultys, `select * from t_faculties limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	return Facultys, nil
}
func (uc *UseCase) GetFaculty(id int) (models.Faculty, error) {
	var FacultyModel models.Faculty
	if err := uc.db.Get(&FacultyModel, `select * from t_faculties where id = $1`, id); err != nil {
		return models.Faculty{}, err
	}
	return FacultyModel, nil
}

func (uc *UseCase) AddFaculty(n *models.Faculty) (models.Faculty, error) {
	var id int

	if n.ID == 0 {
		if err := uc.db.QueryRow(`insert into t_faculties(name) values($1) returning id`, n.Name).Scan(&id); err != nil {
			return models.Faculty{}, err
		}
	} else {
		if err := uc.db.QueryRow(`update t_faculties set name = $1 where id = $2 returning id`, n.Name, n.ID).Scan(&id); err != nil {
			return models.Faculty{}, err
		}
	}

	var oneFaculty models.Faculty
	if err := uc.db.Get(&oneFaculty, `select * from t_faculties where id = $1`, id); err != nil {
		return models.Faculty{}, err
	}

	return oneFaculty, nil
}
func (uc *UseCase) DeleteFaculty(id int) error {
	if _, err := uc.db.Exec(`delete from t_faculties where id = $1`, id); err != nil {
		return err
	}
	return nil
}
