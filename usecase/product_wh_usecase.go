package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type ProductWhUsecase interface {
	FindProducts() any
	FindProductById(id string) any
	Input(newProduct *entity.ProductWh) string
	Edit(product *entity.ProductWh) string
	Output(id string) string
}

type productWhUsecase struct {
	productWhRepo repository.ProductWhRepo
}

func NewProductWhUsecase(productWhRepo repository.ProductWhRepo) ProductWhUsecase {
	return &productWhUsecase{
		productWhRepo: productWhRepo,
	}
}

func (u *productWhUsecase) FindProducts() any {
	return u.productWhRepo.GetAll()
}

func (u *productWhUsecase) FindProductById(id string) any {
	return u.productWhRepo.GetById(id)
}

func (u *productWhUsecase) Input(newProduct *entity.ProductWh) string {
	return u.productWhRepo.Create(newProduct)
}

func (u *productWhUsecase) Edit(product *entity.ProductWh) string {
	return u.productWhRepo.Update(product)
}

func (u *productWhUsecase) Output(id string) string {
	return u.productWhRepo.Delete(id)
}
