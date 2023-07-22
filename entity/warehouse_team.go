package entity

type WarehouseTeam struct {
	ID       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Photo    string `json:"photo"`
}

type EmployeeResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}
