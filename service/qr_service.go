package service

import (
	"adriandidimqttgate/model/web"

	"golang.org/x/net/context"
)

type QrService interface {
	ScanQr(ctx context.Context, request web.ScanQrRequest, userId uint) web.ScanQrResponse
}