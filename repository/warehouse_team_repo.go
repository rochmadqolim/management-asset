package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
)

type WarehouseTeamRepo interface {
	GetAll() any
	GetById(id string) any
	GetByEmail(email string) (*entity.WarehouseTeam, error)
	Create(newEmployee *entity.WarehouseTeam) string
	Update(employee *entity.WarehouseTeam) string
	Delete(id string) string
}

type warehouseTeamRepo struct {
	db *sql.DB
}

func NewWarehouseTeamRepo(db *sql.DB) WarehouseTeamRepo {
	repo := new(warehouseTeamRepo)
	repo.db = db

	return repo
}

func (r *warehouseTeamRepo) GetAll() any {
	var employees []entity.EmployeeResponse

	query := "SELECT id, name, email, phone, photo FROM admin_wh"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee entity.EmployeeResponse

		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone, &employee.Photo); err != nil {
			log.Println(err)
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(employees) == 0 {
		return "no data"
	}

	return employees
}

func (r *warehouseTeamRepo) GetById(id string) any {
	var employeeInDb entity.WarehouseTeam

	query := "SELECT photo FROM admin_wh WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&employeeInDb.Photo)

	if err != nil {
		log.Println(err)
	}

	if employeeInDb.Photo == "" {
		return "employee not found"
	}

	return employeeInDb.Photo
}

func (r *warehouseTeamRepo) GetByEmail(email string) (*entity.WarehouseTeam, error) {
	var employeeInDb entity.WarehouseTeam

	query := "SELECT id, name, email, password, phone, photo FROM admin_wh WHERE email = $1"
	row := r.db.QueryRow(query, email)

	err := row.Scan(&employeeInDb.ID, &employeeInDb.Name, &employeeInDb.Email, &employeeInDb.Password, &employeeInDb.Phone, &employeeInDb.Photo)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}

	return &employeeInDb, nil
}

func (r *warehouseTeamRepo) Create(newEmployee *entity.WarehouseTeam) string {
	query := "INSERT INTO admin_wh(id, name, email, password, phone, photo) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, newEmployee.ID, newEmployee.Name, newEmployee.Email, newEmployee.Password, newEmployee.Phone, newEmployee.Photo)

	if err != nil {
		log.Println(err)
		return "failed to create employee"
	}

	return "New employee created successfully"
}

func (r *warehouseTeamRepo) Update(employee *entity.WarehouseTeam) string {
	res, err := r.GetByEmail(employee.Email) //respon

	// jika tidak ada, return pesan
	if err != nil {
		return err.Error()
	}

	// jika ada maka update user
	query := "UPDATE admin_wh SET id = $1, name = $2, email = $3, password = $4, phone = $5 WHERE email = $6"
	_, err = r.db.Exec(query, employee.ID, employee.Name, employee.Email, employee.Password, employee.Phone, res.Email)

	if err != nil {
		log.Println(err)
		return "failed to update employee"
	}

	return fmt.Sprintf("employee %s updated successfully", res.Name)
}

func (r *warehouseTeamRepo) Delete(id string) string {
	res := r.GetById(id)

	// jika tidak ada, return pesan
	if res == "employee not found" {
		return res.(string)
	}

	// jika ada, delete user
	query := "DELETE FROM admin_wh WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete employee"
	}

	return fmt.Sprintf("employee with id %s deleted successfully", id)
}
