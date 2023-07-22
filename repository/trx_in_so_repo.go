package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
	"time"

	"github.com/tealeg/xlsx"
)

type TrxInSoRepo interface {
	EnrollInsertTrxInSo(EnrollTrxInSt *entity.TrxInSo) string
	EnrollInsertReportConfirm() string
	EnrollInsertReportInterim() string
}

type trxInSoRepo struct {
	db *sql.DB
}

func (trx *trxInSoRepo) EnrollInsertTrxInSo(enrollTrxInSo *entity.TrxInSo) string {
	tx, err := trx.db.Begin()
	if err != nil {
		return "Transaction rollback"
	}

	trx.InserTrxInSo(enrollTrxInSo, tx)
	stock := trx.SumTrxProductSo(enrollTrxInSo.ProductStSoId, tx)
	trx.UpdateProductSo(stock, enrollTrxInSo.ProductStSoId, tx)
	trx.UpdateDifference(tx)
	trx.DeleteTrxInSo(enrollTrxInSo.ProductStSoId, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "product not found"
	} else {
		return "Transaction in stock opname  commited"
	}

}

func (trx *trxInSoRepo) EnrollInsertReportInterim() string {

	tx, err := trx.db.Begin()
	if err != nil {
		return "Transaction rollback"
	}

	total := trx.GetTotalLoss(tx)
	minPN, minTot := trx.GetNameTotalMinLoss(tx)
	maxPN, maxTot := trx.GetNameTotalMaxLoss(tx)

	trx.InsertReportInterimSo(total, minPN, minTot, maxPN, maxTot, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "product not found"
	} else {
		return "Transaction in stock opname  commited"
	}

}

func (trx *trxInSoRepo) EnrollInsertReportConfirm() string {

	tx, err := trx.db.Begin()
	if err != nil {
		return "Transaction rollback"
	}

	total := trx.GetTotalLoss(tx)
	minPN, minTot := trx.GetNameTotalMinLoss(tx)
	maxPN, maxTot := trx.GetNameTotalMaxLoss(tx)

	trx.InsertReportDetailSo(total, minPN, minTot, maxPN, maxTot, tx)
	// trx.ExportExcel(tx)
	trx.ProductStEqualProductSo(tx)
	trx.TruncateSoInterim(tx)
	trx.SetNewProductSo(tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "product not found"
	} else {
		return "Transaction in stock opname  commited"
	}

}

// ========================================================for Report Confirm ======================================================================================================
// Ambil barang dengan minimal loss

func (trx *trxInSoRepo) InsertReportDetailSo(total int, minPN string, minTot int, maxPN string, maxTot int, tx *sql.Tx) {
	query := `INSERT INTO report_so_detail (total_loss, product_min , total_min, product_max,total_max,created_at) values ($1,$2,$3,$4,$5,$6)`

	_, err := tx.Exec(query, total, minPN, minTot, maxPN, maxTot, time.Now())
	Validate(err, "Insert", tx)

}

func (trx *trxInSoRepo) GetTotalLoss(tx *sql.Tx) int {
	query := `select sum(diff_price) from product_so;`

	var totalLoss int

	err := tx.QueryRow(query).Scan(&totalLoss)

	Validate(err, "Select", tx)

	return totalLoss
}

func (trx *trxInSoRepo) GetNameTotalMinLoss(tx *sql.Tx) (string, int) {
	query := `select pst.product_name , pso.diff_price from product_so as pso
	join product_st as pst
	on pso.product_st_id =pst.id
	where diff_price <> 0
	order by diff_price desc limit 1`

	var minProduct string
	minDiffPrice := 0

	err := tx.QueryRow(query).Scan(&minProduct, &minDiffPrice)

	Validate(err, "Select", tx)

	return minProduct, minDiffPrice

}

func (trx *trxInSoRepo) GetNameTotalMaxLoss(tx *sql.Tx) (string, int) {
	query := `select pst.product_name , pso.diff_price from product_so as pso
	join product_st as pst
	on pso.product_st_id =pst.id
	where diff_price <> 0
	order by diff_price asc limit 1`

	var maxProduct string
	maxDiffPrice := 0

	err := tx.QueryRow(query).Scan(&maxProduct, &maxDiffPrice)

	Validate(err, "Select", tx)

	return maxProduct, maxDiffPrice

}
func (trx *trxInSoRepo) SetNewProductSo(tx *sql.Tx) {
	query := `update product_so set stock = 0, diff_price=0, diff_stock = 0`
	_, err := tx.Exec(query)
	Validate(err, "Update", tx)
}

func (trx *trxInSoRepo) TruncateSoInterim(tx *sql.Tx) {
	query := `Truncate table interim_so_report `
	_, err := tx.Exec(query)
	Validate(err, "Truncate", tx)
}
func (trx *trxInSoRepo) ProductStEqualProductSo(tx *sql.Tx) {
	query := `update product_st set stock = product_st.stock +product_so.diff_stock from product_so where product_st.id = product_so.product_st_id;`
	_, err := tx.Exec(query)
	Validate(err, "Update", tx)
}

//================================================================ for eneoll trx so=================================================================================================================

func (trx *trxInSoRepo) InserTrxInSo(enrollTxtInSo *entity.TrxInSo, tx *sql.Tx) {
	query := "INSERT INTO trx_so (product_so_st_id,stock) values($1,$2)"

	_, err := tx.Exec(query, enrollTxtInSo.ProductStSoId, enrollTxtInSo.Stock)
	Validate(err, "Insert", tx)

}

func (trx *trxInSoRepo) SumTrxProductSo(productSoStId string, tx *sql.Tx) int {
	query := `select sum(pso.stock + txs.stock) from trx_so as txs
	join product_so as pso on txs.product_so_st_id = pso.product_st_id
	where txs.product_so_st_id = $1;`
	stock := 0
	err := tx.QueryRow(query, productSoStId).Scan(&stock)

	Validate(err, "Select", tx)
	return stock

}

func (trx *trxInSoRepo) UpdateProductSo(stock int, producSoStId string, tx *sql.Tx) {
	query := `update product_so set stock = $1 where product_st_id = $2`
	_, err := tx.Exec(query, stock, producSoStId)
	Validate(err, "Update", tx)
}

func (trx *trxInSoRepo) UpdateDifference(tx *sql.Tx) {
	query := `update product_so set diff_stock = product_so.stock - product_st.stock, diff_price = (product_st.price * product_so.stock)-(product_st.price * product_st.stock) 
	from product_st
	where product_so.product_st_id = product_st.id;`
	_, err := tx.Exec(query)
	Validate(err, "Update", tx)
}

func (trx *trxInSoRepo) DeleteTrxInSo(productSoStId string, tx *sql.Tx) {
	query := `delete from trx_so where product_so_st_id = $1`
	_, err := tx.Exec(query, productSoStId)
	Validate(err, "Delete", tx)
}
func (trx *trxInSoRepo) InsertReportInterimSo(total int, minPN string, minTot int, maxPN string, maxTot int, tx *sql.Tx) {
	query := `INSERT INTO interim_so_report (total_loss, product_min , total_min, product_max,total_max,created_at) values ($1,$2,$3,$4,$5,$6)`

	waktu := time.Now()

	_, err := tx.Exec(query, total, minPN, minTot, maxPN, maxTot, waktu)
	Validate(err, "Insert", tx)

}

//=============================================================export excel==========================================================================

func (trx *trxInSoRepo) ExportExcel(tx *sql.Tx) {

	rows, err := trx.db.Query("SELECT * FROM product_so")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// buat file Excel
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("Data")
	if err != nil {
		log.Fatal(err)
	}

	// ambil nama kolom dari tabel
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// tambahkan header ke sheet Excel
	row := sheet.AddRow()
	for _, col := range columns {
		cell := row.AddCell()
		cell.Value = col
	}

	// tambahkan data ke sheet Excel
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}
		err := rows.Scan(valuePointers...)
		if err != nil {
			log.Fatal(err)
		}
		row := sheet.AddRow()
		for _, val := range values {
			cell := row.AddCell()
			cell.Value = fmt.Sprintf("%v", val)
		}
	}

	// simpan file Excel ke disk
	err = xlFile.Save("data.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	Validate(err, "Insert", tx)

	log.Println("Data telah berhasil disimpan ke file Excel")
}

func NewTrxInSoRepo(db *sql.DB) TrxInSoRepo {
	repo := new(trxInSoRepo)

	repo.db = db

	return repo
}
