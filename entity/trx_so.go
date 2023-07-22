package entity

type TrxInSo struct {
	ID            int    `json:"id"`
	ProductStSoId string `json:"product_so"`
	Stock         int    `json:"stock"`
}
