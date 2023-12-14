package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse
}