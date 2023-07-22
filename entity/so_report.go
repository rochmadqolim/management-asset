package entity

import (
	"time"
)

type ReportSoRes struct {
	ID         int       `json:"id"`
	TotalLoss  int       `json:"total_loss"`
	ProductMin string    `json:"product_min"`
	TotalMin   int       `json:"total_min"`
	ProductMax string    `json:"product_max"`
	TotalMax   int       `json:"total_max"`
	CreatedAt  time.Time `json:"created_at"`
}
