package auth_srv

import (
	"fmt"

	"google.golang.org/grpc"
	"oliva-back-api-gateway/internal/config"
	"oliva-back-api-gateway/pkg/auth_srv/proto"
)

type ServiceClient struct {
	Client proto.UserClient
}

func InitServiceClient(c *config.Config) proto.UserClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSrv.Host+":"+c.AuthSrv.Port, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewUserClient(cc)
}
