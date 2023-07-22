package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
)

type StoreTeamRepo interface {
	GetAll() any
	GetById(id string) any
	Create(newSeller *entity.Storeteam) string
	Update(seller *entity.Storeteam) string
	Delete(id string) string
}

type storeteamRepo struct {
	db *sql.DB
}

func (r *storeteamRepo) GetAll() any {
	var sellers []entity.Storeteam

	query := "SELECT id, name, email, phone, photo FROM st_team"
	rows, err := r.db.Query(query)
	fmt.Println(rows)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var seller entity.Storeteam

		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Email, &seller.Phone, &seller.Photo); err != nil {
			log.Println(err)
		}

		sellers = append(sellers, seller)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(sellers)
	if len(sellers) == 0 {
		return "no data"
	}
	return sellers
}

func (r *storeteamRepo) GetById(id string) any {
	var seller entity.Storeteam

	query := "SELECT photo FROM st_team WHERE id = $1"

	row := r.db.QueryRow(query, id)

	err := row.Scan(&seller.Photo)
	if err != nil {
		log.Println(err)
	}

	if seller.Photo == "" {
		return "seller not found"
	}

	return seller.Photo
}

func (r *storeteamRepo) Create(newSeller *entity.Storeteam) string {

	query := "INSERT INTO st_team(id, name, email, password, phone, photo) VALUES($1,$2,$3,$4,$5, $6);"

	_, err := r.db.Exec(query, newSeller.ID, newSeller.Name, newSeller.Email, newSeller.Password, newSeller.Phone, newSeller.Photo)
	if err != nil {
		log.Println(err)
		return "failed to create seller"
	}

	return "seller created successfully"

}

func (r *storeteamRepo) Update(seller *entity.Storeteam) string {

	res := r.GetById(seller.ID)
	fmt.Println(seller)
	if res == "seller not found" {
		return res.(string)
	}

	query := "UPDATE st_team SET name=$1, email=$2, password=$3, phone=$4 WHERE id=$5;"
	_, err := r.db.Exec(query, seller.Name, seller.Email, seller.Password, seller.Phone, seller.ID)

	if err != nil {
		log.Println(err)
		return "failed to update seller"
	}

	return fmt.Sprintf("seller with id %s update succesfully", seller.ID)
}

func (r *storeteamRepo) Delete(id string) string {
	res := r.GetById(id)
	if res == "seller not found" {
		return res.(string)
	}

	query := "DELETE FROM st_team WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return "failed to delete seller"
	}

	return fmt.Sprintf("seller with id %s deleted succesfully", id)
}

func NewStoreteamRepo(db *sql.DB) StoreTeamRepo {
	repo := new(storeteamRepo)

	repo.db = db

	return repo
}
