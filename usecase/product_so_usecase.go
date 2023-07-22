package usecase

import (
	"go_inven_ctrl/repository"
)

type ProductSoUsecase interface {
	FindAllProductSo() any
	FindByLessThan(stock int) any
}

type productSoUsecase struct {
	productSoRepo repository.ProductSoRepo
}

func (u *productSoUsecase) FindAllProductSo() any {
	return u.productSoRepo.GetAllProductSo()
}

func (u *productSoUsecase) FindByLessThan(stock int) any {
	return u.productSoRepo.GetByLessThan(stock)
}

func NewProductSoUsecase(productSo repository.ProductSoRepo) ProductSoUsecase {
	return &productSoUsecase{productSoRepo: productSo}
}
