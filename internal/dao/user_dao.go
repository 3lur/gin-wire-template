package dao

import (
	"gin-mini-mall/internal/model"
)

type UserDao interface {
	FindById(id any) (*model.User, error)
	FindAll() ([]*model.User, error)
	Save(user *model.User) error
}

type userDao struct {
	*Dao
}

// FindAll implements UserDao.
// FindAll() returns all users
func (dao userDao) FindAll() ([]*model.User, error) {
	var users []*model.User
	return users, nil
}

// Save implements UserDao.
// Save() save a user
func (dao userDao) Save(user *model.User) error {
	sql := "insert into user(username,password) values(?,?)"
	_, err := dao.db.Exec(sql, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// FindById implements UserDao.
// FindById() returns a user by id
func (dao userDao) FindById(id any) (*model.User, error) {
	var user model.User
	return &user, nil
}

// NewUserDao return  a userDao
func NewUserDao(dao *Dao) UserDao {
	return &userDao{dao}
}
