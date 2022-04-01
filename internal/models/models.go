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
	ID           int    `json:"id" db:"id"`
	Sname        string `json:"sname" db:"sname"`
	Name         string `json:"name" db:"name"`
	Pname        string `json:"pname" db:"pname"`
	Body         string `json:"body" db:"body"`
	IsActive     bool   `json:"is_active" db:"is_active"`
	IsAdmin      bool   `json:"is_admin" db:"is_admin"`
	IsMember     bool   `json:"is_member" db:"is_member"`
	CreateDate   string `json:"create_date" db:"create_date"`
	FileObjectID string `json:"file_object_id" db:"file_object_id"`
	Files        []File `json:"files,omitempty" db:"-" `
}
type File struct {
	ID         int    `json:"id" db:"id"`
	FilePath   string `json:"file_path" db:"file_path"`
	Comment    string `json:"comment,omitempty" db:"comment"`
	IsActive   bool   `json:"is_active,omitempty" db:"is_active"`
	CreateDate string `json:"create_date,omitempty" db:"create_date"`
}
