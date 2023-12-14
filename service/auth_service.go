package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) string
	Register(ctx context.Context, request web.RegisterRequest)
}