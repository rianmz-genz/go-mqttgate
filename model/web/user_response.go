package web

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Office OfficeResponse `json:"office"`
	Role   RoleResponse   `json:"role"`
}
