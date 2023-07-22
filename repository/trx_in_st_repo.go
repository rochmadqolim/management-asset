package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
	"time"
)

type TrxInStRepo interface {
	EnrollInsertTrxInSt(EnrollTrxInSt *entity.TrxInST) string
}

type trxInStRepo struct {
	db *sql.DB
}

func (trx *trxInStRepo) EnrollInsertTrxInSt(EnrollTrxInSt *entity.TrxInST) string {
	tx, err := trx.db.Begin()
	if err != nil {
		return "product not found"
	}
	trx.InserTrxInSt(EnrollTrxInSt, tx)
	name, lastStock := trx.GetProductName(EnrollTrxInSt.ProductStId, tx)

	action := trx.GetAct(EnrollTrxInSt.ProductStId, tx)
	if action == "input" {
		addstock := trx.SumTrxProductSt(EnrollTrxInSt.ProductStId, tx)
		fmt.Println(addstock)
		trx.UpdateProductSt(addstock, EnrollTrxInSt.ProductStId, tx)
	} else if action == "retur" || action == "sold" {
		minstock := trx.MinTrxProductSt(EnrollTrxInSt.ProductStId, tx)
		fmt.Println(minstock)
		trx.UpdateProductSt(minstock, EnrollTrxInSt.ProductStId, tx)
	}
	trx.InserReportInSt(EnrollTrxInSt, name, lastStock, tx)
	trx.DeleteTrxInSt(EnrollTrxInSt.ProductStId, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "Transaction rollback"
	} else {
		return "Transaction in store  commited"
	}

}

func (trx *trxInStRepo) InserTrxInSt(enrollTxtInSt *entity.TrxInST, tx *sql.Tx) {
	query := "INSERT INTO trx_st (id,product_st_id, stock_in,act,created_at) values($1,$2,$3,$4,$5)"

	_, err := tx.Exec(query, &enrollTxtInSt.ID, &enrollTxtInSt.ProductStId, &enrollTxtInSt.StockIn, &enrollTxtInSt.Act, &enrollTxtInSt.CreatedAt)
	Validate(err, "Insert", tx)
	fmt.Println(enrollTxtInSt.Act)
	// return enrollTxtInSt.Act

}

func (trx *trxInStRepo) InserReportInSt(enrollTxtInSt *entity.TrxInST, name string, lastStock int, tx *sql.Tx) {
	query := `INSERT INTO report_trx_st (product_st_id, stock_in,product_name,act,last_stock,created_at) values($1,$2,$3,$4,$5,$6)`
	enrollTxtInSt.CreatedAt = time.Now()
	_, err := tx.Exec(query, enrollTxtInSt.ProductStId, enrollTxtInSt.StockIn, name, enrollTxtInSt.Act, lastStock, enrollTxtInSt.CreatedAt)
	Validate(err, "Insert", tx)

}

func (trx *trxInStRepo) GetProductName(productStId string, tx *sql.Tx) (string, int) {
	fmt.Println("test")
	query := `select pst.product_name, pst.stock from trx_st as tx 
	JOIN product_st as pst on tx.product_st_id =pst.id 
	WHERE product_st_id =$1`
	var name string
	var stock int
	err := tx.QueryRow(query, productStId).Scan(&name, &stock)
	fmt.Println(name)
	Validate(err, "Select", tx)
	return name, stock
}

// Get Category
func (trx *trxInStRepo) GetAct(productStId string, tx *sql.Tx) string {
	query := `select act from trx_st where product_st_id =$1`
	var act string
	err := tx.QueryRow(query, productStId).Scan(&act)
	fmt.Println(act)
	Validate(err, "Select", tx)
	return act
}

func (trx *trxInStRepo) SumTrxProductSt(productStId string, tx *sql.Tx) int {
	query := `SELECT SUM( pst.stock + tx.stock_in ) FROM trx_st as tx 
	JOIN product_st as pst ON tx.product_st_id = pst.id
	WHERE tx.product_st_id = $1
	;
	`
	stock := 0
	err := tx.QueryRow(query, productStId).Scan(&stock)

	Validate(err, "Select", tx)
	return stock

}
func (trx *trxInStRepo) MinTrxProductSt(productStId string, tx *sql.Tx) int {
	query := `SELECT SUM( pst.stock - tx.stock_in) FROM trx_st as tx 
	JOIN product_st as pst ON tx.product_st_id = pst.id
	WHERE tx.product_st_id = $1
	;
	`
	stock := 0
	err := tx.QueryRow(query, productStId).Scan(&stock)

	Validate(err, "Select", tx)
	return stock

}

func (trx *trxInStRepo) UpdateProductSt(stock int, producStId string, tx *sql.Tx) {
	query := `UPDATE product_st SET stock =$1 WHERE id =$2`
	_, err := tx.Exec(query, stock, producStId)
	Validate(err, "Update", tx)
}

func (trx *trxInStRepo) DeleteTrxInSt(productStId string, tx *sql.Tx) {
	query := `delete from trx_st where product_st_id = $1`
	_, err := tx.Exec(query, productStId)
	Validate(err, "Delete", tx)
}

func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("successfully " + message + " data")
	}
}

func NewTrxInStRepo(db *sql.DB) TrxInStRepo {
	repo := new(trxInStRepo)

	repo.db = db

	return repo
}
