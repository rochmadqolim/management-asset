package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type WarehouseTeamUsecase interface {
	FindEmployees() any
	FindEmployeeById(id string) any
	Register(req *entity.WarehouseTeam) string
	Edit(employee *entity.WarehouseTeam) string
	Unreg(id string) string
}

type warehouseTeamUsecase struct {
	warehouseTeamRepo repository.WarehouseTeamRepo
}

func NewWarehouseTeamUsecase(warehouseTeamRepo repository.WarehouseTeamRepo) WarehouseTeamUsecase {
	return &warehouseTeamUsecase{
		warehouseTeamRepo: warehouseTeamRepo,
	}
}

func (u *warehouseTeamUsecase) FindEmployees() any {
	return u.warehouseTeamRepo.GetAll()
}

func (u *warehouseTeamUsecase) FindEmployeeById(id string) any {
	return u.warehouseTeamRepo.GetById(id)
}

func (u *warehouseTeamUsecase) Register(req *entity.WarehouseTeam) string {
	return u.warehouseTeamRepo.Create(req)
}

func (u *warehouseTeamUsecase) Edit(employee *entity.WarehouseTeam) string {
	return u.warehouseTeamRepo.Update(employee)
}

func (u *warehouseTeamUsecase) Unreg(id string) string {
	return u.warehouseTeamRepo.Delete(id)
}
