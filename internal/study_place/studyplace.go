package studyplace

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/models"
)

type UseCase struct {
	db *sqlx.DB
}

func NewStudyPlaceUseCase(db *sqlx.DB) *UseCase {
	return &UseCase{db: db}
}

func (uc *UseCase) GetStudyPlaces(limit, offset int) ([]models.StudyPlace, error) {
	StudyPlace := make([]models.StudyPlace, 0)
	if err := uc.db.Select(&StudyPlace, `select * from t_study_place limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	return StudyPlace, nil
}
func (uc *UseCase) GetStudyPlace(id int) (models.StudyPlace, error) {
	var StudyPlaceModel models.StudyPlace
	if err := uc.db.Get(&StudyPlaceModel, `select * from t_study_place where id = $1`, id); err != nil {
		return models.StudyPlace{}, err
	}
	return StudyPlaceModel, nil
}

func (uc *UseCase) AddStudyPlace(n *models.StudyPlace) (models.StudyPlace, error) {
	var id int

	if n.ID == 0 {
		if err := uc.db.QueryRow(`insert into t_study_place(name) values($1) returning id`, n.Name).Scan(&id); err != nil {
			return models.StudyPlace{}, err
		}
	} else {
		if err := uc.db.QueryRow(`update t_study_place set name = $1 where id = $2 returning id`, n.Name, n.ID).Scan(&id); err != nil {
			return models.StudyPlace{}, err
		}
	}

	var oneStudyPlace models.StudyPlace
	if err := uc.db.Get(&oneStudyPlace, `select * from t_study_place where id = $1`, id); err != nil {
		return models.StudyPlace{}, err
	}

	return oneStudyPlace, nil
}
func (uc *UseCase) DeleteStudyPlace(id int) error {
	if _, err := uc.db.Exec(`delete from t_study_place where id = $1`, id); err != nil {
		return err
	}
	return nil
}
