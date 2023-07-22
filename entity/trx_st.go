package entity

import "time"

type TrxInST struct {
	ID          string    `json:"id"`
	ProductStId string    `json:"product_st_id"`
	StockIn     int       `json:"stock"`
	Act         string    `json:"act"`
	CreatedAt   time.Time `json:"created_at"`
}
