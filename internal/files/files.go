package files

import (
	"github.com/jmoiron/sqlx"
	"lrm-backend/internal/models"
	"os"
	"strings"
	"time"
)

type UseCase struct {
	db *sqlx.DB
}

func NewUseCase(db *sqlx.DB) *UseCase {
	return &UseCase{db: db}
}

func (uc *UseCase) GetFiles(files string) ([]models.File, error) {
	files = strings.ReplaceAll(files, "{", "")
	files = strings.ReplaceAll(files, "}", "")

	filesModels := make([]models.File, 0)
	if err := uc.db.Select(&filesModels, `select id, file_path from t_file_object where id in (`+files+`)`); err != nil {
		return nil, err
	}
	return filesModels, nil
}
func (uc *UseCase) GetUserFile(fileID int) (models.File, error) {
	var file models.File
	if err := uc.db.Get(&file, `select id, file_path from t_file_object where id = $1`, fileID); err != nil {
		return models.File{}, err
	}
	return file, nil
}

func (uc *UseCase) AddFile(filePath string) (int, error) {
	os.Mkdir("photo", 0777)
	var id int
	if err := uc.db.QueryRow(`insert into t_file_object(file_path, comment, is_active,create_date) values($1,$2,$3,$4) returning id`, filePath, "test", true, time.Now()).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
