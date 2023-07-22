package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type ProductStUsecase interface {
	FindAllProductsSt() any
	FindProductStById(id string) any
	RegisterProductSt(newStoreTeam *entity.ProductSt) string
	EditProductSt(student *entity.ProductSt) string
	UnregProductSt(id string) string
}

type productStUsecase struct {
	productStRepo repository.ProductStRepo
}

func (u *productStUsecase) FindAllProductsSt() any {

	return u.productStRepo.GetAllProductSt()
}

func (u *productStUsecase) FindProductStById(id string) any {
	return u.productStRepo.GetByIdProductSt(id)
}

func (u *productStUsecase) RegisterProductSt(newProductSt *entity.ProductSt) string {
	return u.productStRepo.CreateProductStAndSo(newProductSt)
}

func (u *productStUsecase) EditProductSt(productSt *entity.ProductSt) string {
	return u.productStRepo.UpdateProductSt(productSt)
}

func (u *productStUsecase) UnregProductSt(id string) string {
	return u.productStRepo.DeleteProductStAndSo(id)
}

func NewProductStUsecase(productSt repository.ProductStRepo) ProductStUsecase {
	return &productStUsecase{productStRepo: productSt}
}
