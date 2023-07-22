package repository

import (
	"database/sql"
	"log"

	"go_inven_ctrl/entity"
)

type ProductSoRepo interface {
	GetAllProductSo() any
	GetByLessThan(stock int) any
}

type productSoRepo struct {
	db *sql.DB
}

func (r *productSoRepo) GetAllProductSo() any {
	var productsSo []entity.ProductSo

	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where diff_stock <> 0 order by diff_price asc"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var productSo entity.ProductSo

		if err := rows.Scan(&productSo.ID, &productSo.ProductStId, &productSo.Stock, &productSo.DiffPrice, &productSo.DiffStock); err != nil {
			log.Println(err)
		}

		productsSo = append(productsSo, productSo)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(productsSo) == 0 {
		return "no data"
	}

	return productsSo

}

func (r *productSoRepo) GetByLessThan(stock int) any {
	var productsSo []entity.ProductSo

	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where pso.diff_stock < $1 order by diff_price asc"
	rows, err := r.db.Query(query, stock)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var productSo entity.ProductSo

		if err := rows.Scan(&productSo.ID, &productSo.ProductStId, &productSo.Stock, &productSo.DiffStock, &productSo.DiffPrice); err != nil {
			log.Println(err)
		}

		productsSo = append(productsSo, productSo)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(productsSo) == 0 {
		return "no data"
	}

	return productsSo

}

// ============================================================ Cuma Bantuan ====================================================================================

func NewProductSoRepo(db *sql.DB) ProductSoRepo {
	repo := new(productSoRepo)

	repo.db = db

	return repo
}
