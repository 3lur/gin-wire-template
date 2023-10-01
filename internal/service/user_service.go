package service

import (
	"gin-mini-mall/internal/dao"
	"gin-mini-mall/internal/model"
)

type UserService interface {
	GetUserById(id any) (user *model.User, err error)
	CreateUser(user *model.User) error
}

type userService struct {
	*Service
	dao dao.UserDao
}

// CreateUser create user
func (srv userService) CreateUser(user *model.User) error {
	return srv.CreateUser(user)
}

// GetUserById get user by id
func (srv userService) GetUserById(id any) (user *model.User, err error) {
	return srv.dao.FindById(id)
}

func NewUserService(service *Service, dao dao.UserDao) UserService {
	return &userService{service, dao}
}
