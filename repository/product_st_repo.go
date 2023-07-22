package repository

import (
	"database/sql"
	"fmt"
	"log"

	"go_inven_ctrl/entity"
)

type ProductStRepo interface {
	GetAllProductSt() any
	GetByIdProductSt(id string) any
	UpdateProductSt(productSt *entity.ProductSt) string
	DeleteProductStAndSo(id string) string
	CreateProductStAndSo(regPdtStSo *entity.ProductSt) string
}

type productStRepo struct {
	db *sql.DB
}

func (r productStRepo) CreateProductStAndSo(regPdtStSo *entity.ProductSt) string {
	tx, err := r.db.Begin()
	if err != nil {
		return "product already exist"
	}

	r.CreateProductSt(regPdtStSo, tx)
	r.CreateProductSo(regPdtStSo, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "register failed"
	} else {
		return "register product to store and stock opname success"
	}

}

//=============================================== Delete store and stock opnam  products e=====================================

func (r productStRepo) DeleteProductStAndSo(id string) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "product not found"
	}

	r.DeleteProductSo(id, tx)
	r.DeleteProductSt(id, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "Failed to delete product"
	} else {
		return fmt.Sprintf("delete product with id %s from store and stock opname success", id)
	}

}

//===========================================================Get All Product Store ======================================================================

func (r *productStRepo) GetAllProductSt() any {
	fmt.Println("test")
	var productsSt []entity.ProductSt

	query := "SELECT  id,product_name,price,product_category,stock FROM product_st order by id asc"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var productSt entity.ProductSt

		if err := rows.Scan(&productSt.ID, &productSt.ProductName, &productSt.Price, &productSt.ProductCtg, &productSt.Stock); err != nil {
			log.Println(err)
		}

		productsSt = append(productsSt, productSt)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(productsSt)
	if len(productsSt) == 0 {
		return "no data"
	}

	return productsSt

}

//============================================Get By Id Product store =====================================================================

func (r *productStRepo) GetByIdProductSt(id string) any {
	var productSt entity.ProductSt

	query := "SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&productSt.ID, &productSt.ProductName, &productSt.Price, &productSt.ProductCtg, &productSt.Stock); err != nil {
		log.Println(err)
	}

	if productSt.ID == "" {
		return "productSt not found"
	}

	return productSt

}

// ===========================================Register Product Store==========================================================
func (r *productStRepo) CreateProductSt(newProductSt *entity.ProductSt, tx *sql.Tx) {
	query := "INSERT INTO product_st ( id, product_name, product_category, price) VALUES($1,$2,$3,$4)"
	_, err := r.db.Exec(query, newProductSt.ID, newProductSt.ProductName, newProductSt.ProductCtg, newProductSt.Price)

	if err != nil {
		log.Println(err)
	}
	Validate(err, "Insert", tx)
}

func (r *productStRepo) CreateProductSo(newProductSt *entity.ProductSt, tx *sql.Tx) {
	query := "INSERT INTO product_so (product_st_id) VALUES($1)"
	_, err := r.db.Exec(query, newProductSt.ID)

	if err != nil {
		log.Println(err)
	}
	Validate(err, "Insert", tx)

}

//================================================================================================================================

// ========================================Update Product =========================================================================
func (r *productStRepo) UpdateProductSt(productSt *entity.ProductSt) string {
	res := r.GetByIdProductSt(productSt.ID)
	fmt.Println(productSt)
	if res == "productSt not found" {
		return res.(string)
	}

	query := "UPDATE product_st SET product_name = $1, price = $2,product_category = $3 WHERE id = $4 ;"
	_, err := r.db.Exec(query, productSt.ProductName, productSt.Price, productSt.ProductCtg, productSt.ID)

	if err != nil {
		log.Println(err)
		return "failed to update ProductSt"
	}

	return fmt.Sprintf("ProductSt with id %s updated successfully", productSt.ID)
}

// ===================================================Delete Product ================================================================================
func (r *productStRepo) DeleteProductSt(id string, tx *sql.Tx) {
	res := r.GetByIdProductSt(id)
	if res == "productSt not found" {
		return
	}

	query := "DELETE FROM product_st WHERE id =$1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
	}
	Validate(err, "Insert", tx)

}

func (r *productStRepo) DeleteProductSo(id string, tx *sql.Tx) {
	res := r.GetByIdProductSt(id)
	if res == "productSt not found" {
		return
	}
	query := "DELETE FROM product_so WHERE product_st_id =$1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)

	}
	Validate(err, "Delete", tx)

}

// ==========================================================New Repo==================================================================================================
func NewProductStRepo(db *sql.DB) ProductStRepo {
	repo := new(productStRepo)

	repo.db = db

	return repo
}
