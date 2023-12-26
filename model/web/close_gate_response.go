package web

import "time"

type CloseGateResponse struct {
	Name    string    `json:"name"`
	CloseAt time.Time `json:"closeAt"`
}
