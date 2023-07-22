package entity

type ProductSo struct {
	ID          string `json:"id"`
	ProductStId string `json:"product_name"`
	Stock       int    `json:"stock"`
	DiffStock   int    `json:"diff_stock"`
	DiffPrice   int    `json:"diff_price"`
}
