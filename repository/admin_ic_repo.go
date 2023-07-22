package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
)

type AdminIcRepo interface {
	GetAll() any
	GetById(id string) any
	Create(newAdminIc *entity.AdminIc) string
	Update(adminIc *entity.AdminIc) string
	Delete(id string) string
}

type adminIcRepo struct {
	db *sql.DB
}

// ===================================================================Create ============================================
func (a *adminIcRepo) Create(newAdminIc *entity.AdminIc) string {

	query := "INSERT INTO ic_team(id, name, email, phone, photo, password) VALUES ($1, $2, $3, $4, $5, $6)"
	_, exeErr := a.db.Exec(query, newAdminIc.ID, newAdminIc.Name, newAdminIc.Email, newAdminIc.Phone, newAdminIc.Photo, newAdminIc.Password)

	if exeErr != nil {
		log.Println(exeErr)
		return "failed to create user"
	}

	return "user craeted successfully"
}

// ===================================================================Get All ============================================
func (a *adminIcRepo) GetAll() any {
	var adminIcs []entity.AdminIc

	query := "SELECT id, name, email, phone, photo FROM ic_team"
	rows, err := a.db.Query(query)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var adminIc entity.AdminIc

		if err := rows.Scan(&adminIc.ID, &adminIc.Name, &adminIc.Email, &adminIc.Phone, &adminIc.Photo); err != nil {
			log.Println(err)
		}
		adminIcs = append(adminIcs, adminIc)
	}

	if len(adminIcs) == 0 {
		return "no data"
	}
	return adminIcs
}

// ===================================================================Get By Id ============================================
func (a *adminIcRepo) GetById(id string) any {

	var adminIcInDb entity.AdminIc

	query := "SELECT photo FROM ic_team WHERE id = $1"
	row := a.db.QueryRow(query, id)

	err := row.Scan(&adminIcInDb.Photo)

	if err != nil {
		log.Println(err)
	}

	if adminIcInDb.Photo == "" {
		return "admin ic not found"
	}

	return adminIcInDb.Photo

}

// ===================================================================Update ============================================
func (a *adminIcRepo) Update(adminIc *entity.AdminIc) string {
	res := a.GetById(adminIc.ID)

	if res == "admin not found" {
		return res.(string)
	}

	query := "UPDATE ic_team SET name = $1, email = $2, phone = $3, password = $4 WHERE id = $5"
	_, err := a.db.Exec(query, adminIc.Name, adminIc.Email, adminIc.Phone, adminIc.Password, adminIc.ID)

	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("admin ic with id %s updated successfully", adminIc.ID)

}

// ===================================================================Delete ============================================
func (a *adminIcRepo) Delete(id string) string {
	res := a.GetById(id)

	if res == "admin not found" {
		return res.(string)
	}

	query := "DELETE FROM ic_team WHERE id = $1"
	_, err := a.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete admin ic"
	}

	return fmt.Sprintf("admin ic with id %s deleted successfully", id)
}

func NewAdminIcRepo(db *sql.DB) AdminIcRepo {
	repo := new(adminIcRepo)
	repo.db = db
	return repo
}
