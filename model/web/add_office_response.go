package web

type AddOfficeResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Address string `json:"address"`
}
