package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
)

type ProductWhRepo interface {
	GetAll() any
	GetById(id string) any
	Create(newProduct *entity.ProductWh) string
	Update(product *entity.ProductWh) string
	Delete(id string) string
}

type productWhRepo struct {
	db *sql.DB
}

func NewProductWhRepo(db *sql.DB) ProductWhRepo {
	repo := new(productWhRepo)
	repo.db = db

	return repo
}

func (r *productWhRepo) GetAll() any {
	var products []entity.ProductWh

	query := "SELECT id, product_name, price, product_category, stock FROM product_wh"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.ProductWh

		if err := rows.Scan(&product.ID, &product.ProductName, &product.Price, &product.ProductCategory, &product.Stock); err != nil {
			log.Println(err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(products) == 0 {
		return "no data"
	}

	return products
}

func (r *productWhRepo) GetById(id string) any {
	var productInDb entity.ProductWh

	query := "SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&productInDb.ID, &productInDb.ProductName, &productInDb.Price, &productInDb.ProductCategory, &productInDb.Stock)

	if err != nil {
		log.Println(err)
	}

	if productInDb.ID == "" {
		return "product not found"
	}

	return productInDb
}

func (r *productWhRepo) Create(newProduct *entity.ProductWh) string {
	query := "INSERT INTO product_wh(id, product_name, price, product_category, stock) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, newProduct.ID, newProduct.ProductName, newProduct.Price, newProduct.ProductCategory, newProduct.Stock)

	if err != nil {
		log.Println(err)
		return "failed to create product"
	}

	return "new product created successfully"
}

func (r *productWhRepo) Update(product *entity.ProductWh) string {
	res := r.GetById(product.ID) //respon

	if res == "product not found" {
		return res.(string)
	}

	query := "UPDATE product_wh SET id = $1, product_name = $2, price = $3, product_category = $4, stock = $5 WHERE id = $6"
	_, err := r.db.Exec(query, product.ID, product.ProductName, product.Price, product.ProductCategory, product.Stock, product.ID)

	if err != nil {
		log.Println(err)
		return "failed to update product"
	}

	return fmt.Sprintf("product with ID %s updated successfully", product.ID)
}

func (r *productWhRepo) Delete(id string) string {
	res := r.GetById(id)

	// jika tidak ada, return pesan
	if res == "product not found" {
		return res.(string)
	}

	// jika ada, delete user
	query := "DELETE FROM product_wh WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete product"
	}

	return fmt.Sprintf("product with id %s deleted successfully", id)
}
