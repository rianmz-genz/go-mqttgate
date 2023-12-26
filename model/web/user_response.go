package web

type UserResponse struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Office OfficeResponse `json:"office"`
	Role   RoleResponse   `json:"role"`
}
