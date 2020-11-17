package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Golang_Microservices_Course/mvc/services"
	"github.com/Golang_Microservices_Course/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		fmt.Print("Bad request info!\n")
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonErr, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonErr)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		fmt.Print("User doesn't exist!\n")
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		return
	}

	fmt.Print("Returned user: ", user, "\n")
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
