package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"oliva-back-api-gateway/pkg/auth_srv/proto"
)

type LoginUserRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginUser(ctx *gin.Context, c proto.UserClient) {
	body := LoginUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.LoginUser(context.Background(), &proto.LoginRequest{
		Login:    body.Login,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
