package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) web.UserResponse
}
