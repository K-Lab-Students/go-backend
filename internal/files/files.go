package files

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"lrm-backend/internal/models"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
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

func (uc *UseCase) AddFile(file multipart.File) (int, error) {
	var id int
	var fileObject models.File

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	if err := os.Mkdir("photo", 0777); err != nil {
		return 0, err
	}

	fileNameStart := strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(rand.Intn(1000000)) + ".jpg"
	fileName := "./photo/" + fileNameStart
	//todo: Сделать запись файла
	fileObject.FilePath = "/photo/" + fileNameStart
	fileObject.IsActive = true
	if err := ioutil.WriteFile(fileName, fileContents, 0777); err != nil {
		return 0, err
	}

	fileObject.CreateDate = time.Now().Format("2006-01-02")

	if err := uc.db.QueryRow(`insert into t_file_object(file_path, comment, is_active,create_date) values($1,$2,$3,$4)`, fileObject.FilePath, fileObject.Comment, fileObject.IsActive, fileObject.CreateDate).Scan(&id); err != nil {
		return 0, err
	}

	return 0, nil
}
