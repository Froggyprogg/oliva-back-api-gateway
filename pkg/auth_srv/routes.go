package auth_srv

import (
	"github.com/gin-gonic/gin"
	"oliva-back-api-gateway/internal/config"
	"oliva-back-api-gateway/pkg/auth_srv/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.RegisterUser(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.LoginUser(ctx, svc.Client)
}
