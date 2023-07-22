package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type AdminIcUsecase interface {
	FindAdminIc() any
	FindAdminIcById(id string) any
	Register(newAdminIc *entity.AdminIc) string
	Edit(adminIc *entity.AdminIc) string
	Unreg(id string) string
}

type adminIcUsecase struct {
	adminIcRepo repository.AdminIcRepo
}

func (adm *adminIcUsecase) FindAdminIc() any {
	return adm.adminIcRepo.GetAll()
}

func (adm *adminIcUsecase) FindAdminIcById(id string) any {
	return adm.adminIcRepo.GetById(id)
}

func (adm *adminIcUsecase) Register(newAdminIc *entity.AdminIc) string {

	return adm.adminIcRepo.Create(newAdminIc)
}

func (adm *adminIcUsecase) Edit(adminIc *entity.AdminIc) string {
	return adm.adminIcRepo.Update(adminIc)
}

func (adm *adminIcUsecase) Unreg(id string) string {
	return adm.adminIcRepo.Delete(id)
}

func NewAdminIcUsecase(adminIcRepo repository.AdminIcRepo) AdminIcUsecase {
	return &adminIcUsecase{
		adminIcRepo: adminIcRepo,
	}
}
