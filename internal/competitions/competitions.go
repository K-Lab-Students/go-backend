package competitions

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/files"
	"lrm-backend/internal/models"
	"strings"
	"time"
)

type UseCase struct {
	db   *sqlx.DB
	file *files.UseCase
}

func NewCompetitionUseCase(db *sqlx.DB, file *files.UseCase) *UseCase {
	return &UseCase{db: db, file: file}
}

func (uc *UseCase) GetUserCompetitions(competitions string) ([]models.Competition, error) {
	competitions = strings.ReplaceAll(competitions, "{", "")
	competitions = strings.ReplaceAll(competitions, "}", "")

	userCompetitions := make([]models.Competition, 0)
	if err := uc.db.Select(&userCompetitions, `select id, name, color from t_competitions where id in (`+competitions+`)`); err != nil {
		return nil, err
	}
	return userCompetitions, nil
}

func (uc *UseCase) GetCompetitions(limit, offset int) ([]models.Competition, error) {
	Competitions := make([]models.Competition, 0)
	if err := uc.db.Select(&Competitions, `select * from t_competitions limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range Competitions {
		if Competitions[i].FileObjectID != "" {
			filesPathes, err := uc.file.GetFiles(Competitions[i].FileObjectID)
			if err != nil {
				return nil, err
			}
			Competitions[i].Files = filesPathes
		}
	}

	return Competitions, nil
}
func (uc *UseCase) GetCompetition(id int) (models.Competition, error) {
	var CompetitionModel models.Competition
	if err := uc.db.Get(&CompetitionModel, `select * from t_competitions where id = $1`, id); err != nil {
		return models.Competition{}, err
	}
	if CompetitionModel.FileObjectID != "" {
		filesPathes, err := uc.file.GetFiles(CompetitionModel.FileObjectID)
		if err != nil {
			return models.Competition{}, err
		}
		CompetitionModel.Files = filesPathes
	}
	return CompetitionModel, nil
}

func (uc *UseCase) AddCompetition(n *models.Competition) (models.Competition, error) {
	var id int

	n.CreateDate = time.Now().Format("2006-01-02")
	if n.ID == 0 {
		if err := uc.db.QueryRow(`insert into t_competitions(name,body,color,is_active,create_date,file_object_id) values($1,$2,$3,$4,$5,$6) returning id`, n.Name, n.Body, n.Color, n.IsActive, n.CreateDate, n.FileObjectID).Scan(&id); err != nil {
			return models.Competition{}, err
		}
	} else {
		if err := uc.db.QueryRow(`update t_competitions set name = $1, body = $2, is_active = $3 where id = $4 returning id`, n.Name, n.Body, n.IsActive, n.ID).Scan(&id); err != nil {
			return models.Competition{}, err
		}
	}

	var oneCompetitions models.Competition
	if err := uc.db.Get(&oneCompetitions, `select * from t_competitions where id = $1`, id); err != nil {
		return models.Competition{}, err
	}

	return oneCompetitions, nil
}
func (uc *UseCase) DeleteCompetition(id int) error {
	if _, err := uc.db.Exec(`delete from t_competitions where id = $1`, id); err != nil {
		return err
	}
	return nil
}
