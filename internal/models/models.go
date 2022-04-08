package models

type New struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Body         string `json:"body" db:"body"`
	AuthorID     int    `json:"author_id" db:"author_id"`
	IsActive     bool   `json:"is_active" db:"is_active"`
	IsMain       bool   `json:"is_main" db:"is_main"`
	CreateDate   string `json:"create_date" db:"create_date"`
	FileObjectID string `json:"file_object_id" db:"file_object_id"`
	Files        []File `json:"files,omitempty" db:"-" `
}

type Project struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Body         string `json:"body" db:"body"`
	IsActive     bool   `json:"is_active" db:"is_active"`
	AuthorID     int    `json:"author_id" db:"author_id"`
	CreateDate   string `json:"create_date" db:"create_date"`
	FileObjectID string `json:"file_object_id" db:"file_object_id"`
	Files        []File `json:"files,omitempty" db:"-" `
}
type User struct {
	ID             int     `json:"id" db:"id"`
	Sname          *string `json:"sname" db:"sname"`
	Name           *string `json:"name" db:"name"`
	Pname          *string `json:"pname" db:"pname"`
	Email          string  `json:"email" db:"email"`
	Password       string  `json:"password" db:"password"`
	IsActive       bool    `json:"is_active" db:"is_active"`
	IsAdmin        bool    `json:"is_admin" db:"is_admin"`
	IsMember       bool    `json:"is_member" db:"is_member"`
	CreateDate     *string `json:"create_date" db:"create_date"`
	FileObjectID   *int    `json:"file_object_id" db:"file_object_id"`
	Birthday       *string `json:"birthday,omitempty" db:"birthday" `
	CompetitionsID *string `json:"competitions_id,omitempty" db:"competitions_id" `
	FacultyID      *string `json:"faculty_id,omitempty" db:"faculty_id" `
	FacultyName    *string `json:"faculty_name,omitempty" db:"faculty_name" `
	StudyPlaceID   *string `json:"study_place_id,omitempty" db:"study_place_id" `
	StudyPlaceName *string `json:"study_place_name,omitempty" db:"study_place_name" `

	UserRoleID   int           `json:"user_role_id" db:"-"`
	Files        []File        `json:"files,omitempty" db:"-" `
	Competitions []Competition `json:"competitions,omitempty" db:"-" `
}

type Faculty struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type StudyPlace struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UserUpdateProfile struct {
	ID             int     `json:"id" db:"id"`
	FacultyID      int     `json:"faculty_id" db:"faculty_id"`
	StudyPlaceID   int     `json:"study_place_id" db:"study_place_id"`
	CompetitionsID *string `json:"competitions_id" db:"competitions_id"`
	Sname          string  `json:"sname" db:"sname"`
	Name           string  `json:"name" db:"name"`
	Pname          string  `json:"pname" db:"pname"`
	Birthday       string  `json:"birthday" db:"birthday"`
	FileID         *int    `json:"file_id" db:"file_id"`
	File           *string `json:"file" db:"file"`
}

type Competition struct {
	ID           int    `json:"id,omitempty" db:"id"`
	Name         string `json:"name,omitempty" db:"name"`
	Body         string `json:"body,omitempty" db:"body"`
	Color        string `json:"color,omitempty" db:"color"`
	IsActive     bool   `json:"is_active,omitempty" db:"is_active"`
	CreateDate   string `json:"create_date,omitempty" db:"create_date"`
	FileObjectID string `json:"file_object_id,omitempty" db:"file_object_id"`
	Files        []File `json:"files,omitempty,omitempty" db:"-" `
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Md5 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	IsMember bool   `json:"is_member"`
}
type File struct {
	ID         int    `json:"id" db:"id"`
	FilePath   string `json:"file_path" db:"file_path"`
	Comment    string `json:"comment,omitempty" db:"comment"`
	IsActive   bool   `json:"is_active,omitempty" db:"is_active"`
	CreateDate string `json:"create_date,omitempty" db:"create_date"`
}
