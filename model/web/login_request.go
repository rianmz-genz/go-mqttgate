package web

type LoginRequest struct {
	Email    string `validate:"required,min=1,max=250" json:"email"`
	Password string `validate:"required,min=1,max=250" json:"password"`
}
