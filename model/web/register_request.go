package web

type RegisterRequest struct {
	Name     string `validate:"required,min=1,max=250" json:"name"`
	Email    string `validate:"required,min=1,max=250" json:"email"`
	Password string `validate:"required,min=1,max=250" json:"password"`
	OfficeID uint `validate:"required,min=1,max=250" json:"officeId"`
}
