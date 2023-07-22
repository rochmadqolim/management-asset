package entity

type AdminIc struct {
	ID       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Photo    string `json:"photo"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}
