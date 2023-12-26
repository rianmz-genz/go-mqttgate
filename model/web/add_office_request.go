package web

type AddOfficeRequest struct {
	Name    string `validate:"required,min=1,max=250" json:"name"`
	Code    string `validate:"required,min=1,max=6" json:"code"`
	Address string `validate:"required,min=1,max=250" json:"address"`
}
