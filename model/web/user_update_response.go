package web

type UserUpdateResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	OfficeID uint   `json:"officeID"`
}
