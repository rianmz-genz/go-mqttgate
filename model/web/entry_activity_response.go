package web

import "time"

type EntryActivityResponse struct {
	Name    string    `json:"name"`
	EntryAt time.Time `json:"entryAt"`
}
