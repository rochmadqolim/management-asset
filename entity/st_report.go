package entity

type ReportTrxSt struct {
	// ID          string
	ProductStId string `json:"product_st_id"`
	StockIn     int    `json:"stock"`
	ProductName string `json:"product_name"`
	Act         string `json:"act"`
	LastStock   int    `json:"last_stock"`
	CreatedAt   string `json:"created_at"`
}
