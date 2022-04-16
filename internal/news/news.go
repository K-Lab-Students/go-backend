package news

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/files"
	"lrm-backend/internal/models"
	"os"
	"strings"
	"time"
)

type UseCase struct {
	db   *sqlx.DB
	file *files.UseCase
}

func NewUseCase(db *sqlx.DB, file *files.UseCase) *UseCase {
	return &UseCase{db: db, file: file}
}

func (uc *UseCase) GetNews(limit, offset int) ([]models.New, error) {
	news := make([]models.New, 0)
	if err := uc.db.Select(&news, `select t_news.*,tu.name as author_name from t_news
												join t_user tu on t_news.author_id = tu.id
												order by create_date desc,id desc limit $1 
												offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range news {
		if news[i].FileObjectID != "" {
			filesPathes, err := uc.file.GetFiles(news[i].FileObjectID)
			if err != nil {
				return nil, err
			}
			news[i].Files = filesPathes
		}
	}

	return news, nil
}
func (uc *UseCase) GetNew(id int) (models.New, error) {
	var newModel models.New
	if err := uc.db.Get(&newModel, `select t_news.*,tu.name as author_name from t_news
												join t_user tu on t_news.author_id = tu.id
												where t_news.id = $1`, id); err != nil {
		return models.New{}, err
	}

	if newModel.FileObjectID != "" {
		filesPathes, err := uc.file.GetFiles(newModel.FileObjectID)
		if err != nil {
			return models.New{}, err
		}
		newModel.Files = filesPathes
	}

	return newModel, nil
}

func (uc *UseCase) AddNew(n *models.New) (models.New, error) {
	var id int

	n.CreateDate = time.Now().Format("2006-01-02")
	if n.FileObjectID == "" || n.FileObjectID == "{}" {
		n.FileObjectID = "{1}"
	}
	if n.ID == 0 {
		if err := uc.db.QueryRow(`insert into t_news(name,body,author_id,is_active,is_main,create_date,file_object_id) values($1,$2,$3,$4,$5,$6,$7) returning id`, n.Name, n.Body, n.AuthorID, n.IsActive, n.IsMain, n.CreateDate, n.FileObjectID).Scan(&id); err != nil {
			return models.New{}, err
		}
	} else {
		if err := uc.db.QueryRow(`update t_news set name = $1, body = $2,author_id = $3, is_active = $4,is_main = $5, file_object_id = $6 where id = $7 returning id`, n.Name, n.Body, n.AuthorID, n.IsActive, n.IsMain, n.FileObjectID, n.ID).Scan(&id); err != nil {
			return models.New{}, err
		}
	}

	var oneNews models.New
	if err := uc.db.Get(&oneNews, `select * from t_news where id = $1`, id); err != nil {
		return models.New{}, err
	}

	return oneNews, nil
}
func (uc *UseCase) DeleteNew(id int) error {
	if _, err := uc.db.Exec(`delete from t_news where id = $1`, id); err != nil {
		return err
	}
	return nil
}
func (uc *UseCase) DeleteFiles(id int, dontDeleteFiles []string) error {
	var filesArr string
	if err := uc.db.Get(&filesArr, `select file_object_id from t_news where id = $1`, id); err != nil {
		return err
	}

	filesArr = strings.ReplaceAll(filesArr, "{", "")
	filesArr = strings.ReplaceAll(filesArr, "}", "")

	var f []models.File
	if err := uc.db.Select(&f, `select * from t_file_object where id in (`+filesArr+`)`); err != nil {
		return err
	}

	for i := range f {
		breakFlag := false
		for _, ddf := range dontDeleteFiles {
			if f[i].FilePath == ddf {
				breakFlag = true
				break
			}
		}

		if breakFlag {
			continue
		}

		if err := os.Remove("." + f[i].FilePath); err != nil {
			return err
		}
		if _, err := uc.db.Exec(`delete from t_file_object where id = $1`, f[i].ID); err != nil {
			return err
		}

	}

	return nil
}
