package users

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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
		if *Users[i].FileObjectID != "" {
			filesPathes, err := uc.file.GetFiles(*Users[i].FileObjectID)
			if err != nil {
				return nil, err
			}
			Users[i].Files = filesPathes
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

	var l Login
	fmt.Println("aa:", a.Email)
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
	if err := uc.db.Get(&user, `select * from t_user where id = $1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}
