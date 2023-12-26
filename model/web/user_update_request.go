package web

type UserUpdateRequest struct {
	Name     string `validate:"required,min=1,max=250" json:"name"`
	Email    string `validate:"required,min=1,max=250" json:"email"`
	OfficeID uint   `validate:"required,min=1,max=250" json:"officeId"`
}
