package request

import "gin-mini-mall/internal/model"

type UserCreateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ToUser 将请求体转换为模型
func (r *UserCreateRequest) ToUser() *model.User {
	return &model.User{
		Username: r.UserName,
		Password: r.Password,
	}
}
