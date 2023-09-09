package service

import (
	"errors"

	"github.com/MohamedAbdelrazeq/gin-video-server/entity"
)

type UserService interface {
	LoginUser(entity.LoginUser) (entity.UserToken, error)
	RegisterUser(entity.UserData) (entity.UserData, error)
}

type userService struct {
	usersData []entity.UserData
}

func NewUserService() UserService {
	return &userService{
		usersData: []entity.UserData{},
	}
}

func (service *userService) LoginUser(user entity.LoginUser) (entity.UserToken, error) {
	for _, u := range service.usersData {
		if u.Email == user.Email && u.Passoword == user.Passoword {
			return entity.UserToken{
				Token:        "token",
				RefreshToken: "refreshToken",
			}, nil
		}
	}
	return entity.UserToken{}, errors.New("invalid email or password")
}

func (service *userService) RegisterUser(userData entity.UserData) (entity.UserData, error) {
	service.usersData = append(service.usersData, userData)
	return userData, nil
}
