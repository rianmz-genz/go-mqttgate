package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type SessionService interface {
	GetSessionById(ctx context.Context, sessionId uint) web.SessionResponse
}
