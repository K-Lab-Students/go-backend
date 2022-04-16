package users

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"lrm-backend/internal/competitions"
	Facultys "lrm-backend/internal/faculties"
	"lrm-backend/internal/files"
	"lrm-backend/internal/models"
	"os"
)

type UseCase struct {
	db           *sqlx.DB
	file         *files.UseCase
	competitions *competitions.UseCase
	faculties    *Facultys.UseCase
}

const query = `select t_user.*, tf.name as faculty_name, tsp.name as study_place_name from t_user
         			left join t_faculties tf on t_user.faculty_id = tf.id
         			left join t_study_place tsp on t_user.study_place_id = tsp.id `

func NewUsersUseCase(db *sqlx.DB, file *files.UseCase, competitions *competitions.UseCase, faculties *Facultys.UseCase) *UseCase {
	return &UseCase{
		db:           db,
		file:         file,
		competitions: competitions,
		faculties:    faculties,
	}
}

func (uc *UseCase) GetUsers(limit, offset int) ([]models.User, error) {
	Users := make([]models.User, 0)
	if err := uc.db.Select(&Users, query+` limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range Users {
		if err := uc.getUsersFile(&Users[i]); err != nil {
			return nil, err
		}
		if err := uc.getUsersCompetitions(&Users[i]); err != nil {
			return nil, err
		}
	}

	return Users, nil
}

func (uc *UseCase) Registration(a *models.Auth) (models.User, error) {
	var userID int

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	if err := uc.db.Get(&userID, `insert into t_user(email,password) VALUES ($1,$2) returning id`, a.Email, hashedPassword); err != nil {
		return models.User{}, err
	}

	return uc.GetUserByID(userID)
}

type Login struct {
	Id       int    `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
}

func (uc *UseCase) Login(a *models.Auth) (models.User, error) {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return models.User{}, err
	//}

	fmt.Println(a.Email)
	var l Login
	if err := uc.db.Get(&l, `select id, password from t_user where email LIKE $1`, a.Email); err != nil {
		return models.User{}, errors.New("Ошибка получения данных пользователя")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(l.Password), []byte(a.Password)); err != nil {
		return models.User{}, errors.New("Указан неверный пароль")
	}

	return uc.GetUserByID(l.Id)
}

func (uc *UseCase) GetUserByID(id int) (models.User, error) {
	var user models.User
	if err := uc.db.Get(&user, query+` where t_user.id = $1`, id); err != nil {
		return models.User{}, err
	}

	if err := uc.getUsersFile(&user); err != nil {
		return models.User{}, err
	}
	if err := uc.getUsersCompetitions(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (uc *UseCase) UpdateUserProfile(p *models.UserUpdateProfile) error {
	q := `update t_user set sname = $1, name = $2, pname = $3, birthday = $4, file_object_id = $5, faculty_id = $6, study_place_id = $7, competitions_id = $8 where id = $9`
	if _, err := uc.db.Exec(q, p.Sname, p.Name, p.Pname, p.Birthday, p.FileID, p.FacultyID, p.StudyPlaceID, p.CompetitionsID, p.ID); err != nil {
		return err
	}
	return nil
}
func (uc *UseCase) getUsersFile(user *models.User) error {
	if user.FileObjectID != nil && *user.FileObjectID != 0 {
		filesPathes, err := uc.file.GetUserFile(*user.FileObjectID)
		if err != nil {
			return err
		}
		user.Files = []models.File{filesPathes}
	}
	return nil
}
func (uc *UseCase) getUsersCompetitions(user *models.User) error {
	if user.CompetitionsID != nil && *user.CompetitionsID != "" {
		userCompetitions, err := uc.competitions.GetUserCompetitions(*user.CompetitionsID)
		if err != nil {
			return err
		}
		user.Competitions = userCompetitions
	}
	return nil
}
func (uc *UseCase) SaveUserFile(filePath string) (int, error) {
	id, err := uc.file.AddFile(filePath)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *UseCase) SearchFile(filePath string) (int, error) {
	var fileID int
	if err := uc.db.Get(&fileID, `select id from t_file_object where file_path LIKE $1`, filePath); err != nil {
		if err != nil {
			return 0, err
		}
	}

	return fileID, nil
}

func (uc *UseCase) DeleteFiles(id int) error {
	var fileId int
	if err := uc.db.Get(&fileId, `select file_object_id from t_user where id = $1`, id); err != nil {
		return err
	}

	var f models.File
	if err := uc.db.Get(&f, `select * from t_file_object where id = $1`, fileId); err != nil {
		return err
	}

	if err := os.Remove("." + f.FilePath); err != nil {
		return err
	}

	if _, err := uc.db.Exec(`delete from t_file_object where id = $1`, fileId); err != nil {
		return err
	}

	return nil
}
