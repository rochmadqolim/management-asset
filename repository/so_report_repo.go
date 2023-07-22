package repository

import (
	"database/sql"
	"go_inven_ctrl/entity"
	"log"
)

type ReportSoRepo interface {
	GetAllInterimSoReport() any
	GetAlDetailSoReport() any
}

type reportSoRepo struct {
	db *sql.DB
}

func (r *reportSoRepo) GetAllInterimSoReport() any {
	var reportsSo []entity.ReportSoRes

	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from interim_so_report order by id"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportSo entity.ReportSoRes

		if err := rows.Scan(&reportSo.ID, &reportSo.TotalLoss, &reportSo.ProductMin, &reportSo.TotalMin, &reportSo.ProductMax, &reportSo.TotalMax, &reportSo.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsSo = append(reportsSo, reportSo)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsSo) == 0 {
		return "no data"
	}

	return reportsSo

}

func (r *reportSoRepo) GetAlDetailSoReport() any {
	var reportsSo []entity.ReportSoRes

	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from report_so_detail order by created_at"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportSo entity.ReportSoRes

		if err := rows.Scan(&reportSo.ID, &reportSo.TotalLoss, &reportSo.ProductMin, &reportSo.TotalMin, &reportSo.ProductMax, &reportSo.TotalMax, &reportSo.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsSo = append(reportsSo, reportSo)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsSo) == 0 {
		return "no data"
	}

	return reportsSo

}

func NewReportSoRepo(db *sql.DB) ReportSoRepo {
	repo := new(reportSoRepo)

	repo.db = db

	return repo
}
