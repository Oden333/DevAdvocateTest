package models

type User struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"name" binding:"required" db:"name"`
	Surname    string `json:"surname" binding:"required" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}
