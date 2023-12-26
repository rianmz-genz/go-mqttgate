package service

import (
	"adriandidimqttgate/model/web"
	"context"
)

type OfficeService interface {
	GetEntryActivities(ctx context.Context, sessionId uint, officeId uint) ([]web.EntryActivityResponse, error)
	CloseGate(ctx context.Context, sessionId uint, officeId uint) (web.CloseGateResponse, error)
}
