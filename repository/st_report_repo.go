package repository

import (
	"database/sql"
	"log"

	"go_inven_ctrl/entity"
)

type ReportTrxStRepo interface {
	GetAllReportTrxSt() any
	GetByReportTrxProductStId(id string) any
	GetByDateReportTrxSt(date string) any
}

type reportTrxStRepo struct {
	db *sql.DB
}

func (r *reportTrxStRepo) GetAllReportTrxSt() any {

	var reportsTrxSt []entity.ReportTrxSt

	query := `select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st order by act asc, created_at asc`
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxSt entity.ReportTrxSt

		if err := rows.Scan(&reportTrxSt.ProductStId, &reportTrxSt.StockIn, &reportTrxSt.ProductName, &reportTrxSt.Act, &reportTrxSt.LastStock, &reportTrxSt.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxSt = append(reportsTrxSt, reportTrxSt)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxSt) == 0 {
		return "no data"
	}

	return reportsTrxSt

}

func (r *reportTrxStRepo) GetByReportTrxProductStId(id string) any {
	var reportsTrxSt []entity.ReportTrxSt

	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where product_st_id = $1 order by act asc, created_at asc"

	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxSt entity.ReportTrxSt

		if err := rows.Scan(&reportTrxSt.ProductStId, &reportTrxSt.StockIn, &reportTrxSt.ProductName, &reportTrxSt.Act, &reportTrxSt.LastStock, &reportTrxSt.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxSt = append(reportsTrxSt, reportTrxSt)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxSt) == 0 {
		return "no data or product id not found"
	}

	return reportsTrxSt

}

func (r *reportTrxStRepo) GetByDateReportTrxSt(date string) any {
	var reportsTrxSt []entity.ReportTrxSt
	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where created_at = $1 order by act asc, created_at asc"

	rows, err := r.db.Query(query, date)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxSt entity.ReportTrxSt

		if err := rows.Scan(&reportTrxSt.ProductStId, &reportTrxSt.StockIn, &reportTrxSt.ProductName, &reportTrxSt.Act, &reportTrxSt.LastStock, &reportTrxSt.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxSt = append(reportsTrxSt, reportTrxSt)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxSt) == 0 {
		return "no data or date product not found"
	}

	return reportsTrxSt
}

func NewReportTrxStRepo(db *sql.DB) ReportTrxStRepo {
	repo := new(reportTrxStRepo)

	repo.db = db

	return repo
}
