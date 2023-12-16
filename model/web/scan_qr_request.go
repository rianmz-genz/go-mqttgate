package web

type ScanQrRequest struct {
	Code string `validate:"required,min=1,max=6" json:"code"`
}
