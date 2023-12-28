package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) web.UserResponse
	GetUsersByOfficeId(ctx context.Context, sessionId uint, officeId uint) []web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest, sessionId uint, userId uint) (web.UserUpdateResponse, error)
	Delete(ctx context.Context, sessionId uint, userId uint)
}
