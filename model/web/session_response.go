package web

type SessionResponse struct {
	SessionID uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserId    uint   `json:"userId"`
}
