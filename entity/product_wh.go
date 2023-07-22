package entity

type ProductWh struct {
	ID              string `json:"id" validate:"required"`
	ProductName     string `json:"product_name" validate:"required,min=5"`
	Price           int    `json:"price" validate:"required,min=100"`
	ProductCategory string `json:"product_ctg" validate:"required,min=3"`
	Stock           int    `json:"stock"`
}
