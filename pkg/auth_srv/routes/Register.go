package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"oliva-back-api-gateway/pkg/auth_srv/proto"
)

type RegisterUserRequestBody struct {
	Surname     string `json:"surname"`
	Name        string `json:"name"`
	Middlename  string `json:"middlename"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
}

func RegisterUser(ctx *gin.Context, c proto.UserClient) {
	body := RegisterUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateUser(context.Background(), &proto.PostUserRequest{
		Surname:     body.Surname,
		Name:        body.Name,
		Middlename:  body.Middlename,
		PhoneNumber: body.PhoneNumber,
		Email:       body.Email,
		Sex:         body.Sex,
		Password:    body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
