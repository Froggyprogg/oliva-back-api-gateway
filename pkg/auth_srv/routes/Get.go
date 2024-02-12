package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"oliva-back-api-gateway/pkg/auth_srv/proto"
)

type GetUserRequestBody struct {
	Id_user uint32 `json:"id_user"`
}

func GetUser(ctx *gin.Context, c proto.UserClient) {
	body := GetUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetUser(context.Background(), &proto.GetUserRequest{
		IdUser: body.Id_user,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
