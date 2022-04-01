package projects

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/files"
	"lrm-backend/internal/models"
	"time"
)

type UseCase struct {
	db   *sqlx.DB
	file *files.UseCase
}

func NewProjectUseCase(db *sqlx.DB, file *files.UseCase) *UseCase {
	return &UseCase{db: db, file: file}
}

func (uc *UseCase) GetProjects(limit, offset int) ([]models.Project, error) {
	Projects := make([]models.Project, 0)
	if err := uc.db.Select(&Projects, `select * from t_projects limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range Projects {
		if Projects[i].FileObjectID != "" {
			filesPathes, err := uc.file.GetFiles(Projects[i].FileObjectID)
			if err != nil {
				return nil, err
			}
			Projects[i].Files = filesPathes
		}
	}

	return Projects, nil
}
func (uc *UseCase) GetProject(id int) (models.Project, error) {
	var ProjectModel models.Project
	if err := uc.db.Get(&ProjectModel, `select * from t_projects where id = $1`, id); err != nil {
		return models.Project{}, err
	}
	if ProjectModel.FileObjectID != "" {
		filesPathes, err := uc.file.GetFiles(ProjectModel.FileObjectID)
		if err != nil {
			return models.Project{}, err
		}
		ProjectModel.Files = filesPathes
	}
	return ProjectModel, nil
}

func (uc *UseCase) AddProject(n *models.Project) (models.Project, error) {
	var id int

	n.CreateDate = time.Now().Format("2006-01-02")
	if n.ID == 0 {
		if err := uc.db.QueryRow(`insert into t_projects(name,body,is_active,create_date,author_id,file_object_id) values($1,$2,$3,$4,$5,$6) returning id`, n.Name, n.Body, n.IsActive, n.CreateDate, n.AuthorID, n.FileObjectID).Scan(&id); err != nil {
			return models.Project{}, err
		}
	} else {
		if err := uc.db.QueryRow(`update t_projects set name = $1, body = $2, is_active = $3 where id = $4 returning id`, n.Name, n.Body, n.IsActive, n.ID).Scan(&id); err != nil {
			return models.Project{}, err
		}
	}

	var oneProjects models.Project
	if err := uc.db.Get(&oneProjects, `select * from t_projects where id = $1`, id); err != nil {
		return models.Project{}, err
	}

	return oneProjects, nil
}
func (uc *UseCase) DeleteProject(id int) error {
	if _, err := uc.db.Exec(`delete from t_projects where id = $1`, id); err != nil {
		return err
	}
	return nil
}
