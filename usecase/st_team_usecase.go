package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type StoreteamUsecase interface {
	FindSellers() any
	FindSellerById(id string) any
	Register(newSeller *entity.Storeteam) string
	Edit(seller *entity.Storeteam) string
	Unreg(id string) string
}

type storeteamUsecase struct {
	storeteamRepo repository.StoreTeamRepo
}

func (u *storeteamUsecase) FindSellers() any {

	return u.storeteamRepo.GetAll()
}

func (u *storeteamUsecase) FindSellerById(id string) any {
	return u.storeteamRepo.GetById(id)
}

func (u *storeteamUsecase) Register(newSeller *entity.Storeteam) string {
	return u.storeteamRepo.Create(newSeller)
}

func (u *storeteamUsecase) Edit(seller *entity.Storeteam) string {
	return u.storeteamRepo.Update(seller)
}

func (u *storeteamUsecase) Unreg(id string) string {
	return u.storeteamRepo.Delete(id)
}

func NewStoreteamUsecase(storeteamRepo repository.StoreTeamRepo) StoreteamUsecase {
	return &storeteamUsecase{
		storeteamRepo: storeteamRepo,
	}
}
