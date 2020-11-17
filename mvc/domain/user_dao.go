package domain

import (
	"fmt"
	"net/http"

	"github.com/Golang_Microservices_Course/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Luke", LastName: "Tchang", Email: "ltchang@stanford.edu"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found.", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
