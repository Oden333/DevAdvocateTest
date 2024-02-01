package service

import (
	"DATest/models"
	"DATest/pkg/handler/helper"
	"DATest/pkg/repository"
)

type User_service struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *User_service {
	return &User_service{repo: repo}
}

func (s *User_service) CreateUser(user models.User) (int, error) {
	//Здесь можно добавить валидацию user
	return s.repo.CreateUser(user)
}

func (s *User_service) GetAllUsers(limit int, offset int) (int, []models.User, error) {

	return s.repo.GetAllUsers(limit, offset)

}

func (s *User_service) GetCertainUsers(limit int, offset int, filter map[string]string) (int, []models.User, error) {

	//Добавить валидацию filter?

	return s.repo.GetCertainUsers(limit, offset, filter)
}

func (s *User_service) DeleteUser(userId int) error {

	return s.repo.DeleteUser(userId)
}

func (s *User_service) UpdateUser(userId int, user helper.UserData) error {
	//Проверяем, есть ли юзер с таким Id в бд
	return s.repo.UpdateUser(userId, user)
	//Если такого нет, то создаём нового
}
