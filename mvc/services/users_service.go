package services

import (
	"github.com/Golang_Microservices_Course/mvc/domain"
	"github.com/Golang_Microservices_Course/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
