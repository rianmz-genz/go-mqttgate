package web

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	OfficeID uint   `json:"officeId"`
}
