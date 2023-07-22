package usecase

import (
	"fmt"
	"go_inven_ctrl/entity"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// =============== dummy data ====================
var dummyProductSt = []entity.ProductSt{
	{
		ID: "pst1", ProductName: "Oreo", ProductCtg: "makanan", Price: 5000,
	},
	{
		ID: "pst2", ProductName: "Richeese", ProductCtg: "makanan", Price: 3000,
	},
	{
		ID: "pst3", ProductName: "Teh Kotak", ProductCtg: "minuman", Price: 4000,
	},
}

// ==============================================
type productStrepoMock struct {
	mock.Mock
}

func (r *productStrepoMock) CreateProductStAndSo(regPdtStSo *entity.ProductSt) string {
	args := r.Called(regPdtStSo)
	if args[0] == "register product to store and stock opname success" {
		return "register product to store and stock opname success"
	}
	if args[0] == "product already exist" {
		return "product already exist"
	}
	return "register failed"
}
func (r *productStrepoMock) UpdateProductSt(productSt *entity.ProductSt) string {
	args := r.Called(productSt)
	if args[0] == fmt.Sprintf("ProductSt with id %s updated successfully", dummyProductSt[0].ID) {
		return fmt.Sprintf("ProductSt with id %s updated successfully", productSt.ID)
	}
	if args[0] == "product not found" {
		return "product not found"
	}
	return "failed to update Product"
}
func (r *productStrepoMock) GetByIdProductSt(id string) any {
	args := r.Called(id)
	if args[0] == id {
		for _, product := range dummyProductSt {
			if product.ID == args[0] {
				return product
			}
		}
	}
	return "productSt not found"
}

func (r *productStrepoMock) GetAllProductSt() any {
	args := r.Called()
	if args[0] == nil {
		return "no data"
	}
	return dummyProductSt
}

func (r *productStrepoMock) DeleteProductStAndSo(id string) string {
	args := r.Called(id)
	for _, product := range dummyProductSt {
		if args[0] == product.ID {
			return fmt.Sprintf("delete product with id %s from store and stock opname success", id)
		} else if args[0] == "" {
			return "Failed to delete product"
		}
	}

	return "product not found"
}

// =====================================================================
type ProductStUCTestSuite struct {
	productStrepoMock *productStrepoMock
	suite.Suite
}

// ====================function Test Success and Failed==============

// ======== UnReg =========
func (suite *ProductStUCTestSuite) TestUnRegProductStAndSo_Success() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	id := "pst1"
	suite.productStrepoMock.On("DeleteProductStAndSo", id).Return(id)
	value := ProductSt.UnregProductSt(id)
	assert.Equal(suite.T(), value, fmt.Sprintf("delete product with id %s from store and stock opname success", id))
}

func (suite *ProductStUCTestSuite) TestUnRegProductStAndSo_Failed() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	id := ""
	suite.productStrepoMock.On("DeleteProductStAndSo", id).Return(id)
	value := ProductSt.UnregProductSt(id)
	assert.Equal(suite.T(), value, "Failed to delete product")
}

func (suite *ProductStUCTestSuite) TestUnRegProductStAndSo_Failed2() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	id := "pst4"
	suite.productStrepoMock.On("DeleteProductStAndSo", id).Return(id)
	value := ProductSt.UnregProductSt(id)
	assert.Equal(suite.T(), value, "product not found")
}

// ========= Find All ============
func (suite *ProductStUCTestSuite) TestFindAllProductSt_success() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("GetAllProductSt").Return(dummyProductSt)
	value := ProductSt.FindAllProductsSt()
	assert.Equal(suite.T(), value, dummyProductSt)
}
func (suite *ProductStUCTestSuite) TestFindAllProductSt_Failed() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("GetAllProductSt").Return(nil)
	value := ProductSt.FindAllProductsSt()
	assert.Equal(suite.T(), "no data", value)
}

// ========= Find By Id =========
func (suite *ProductStUCTestSuite) TestFindIdProductSt_success() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	id := "pst1"
	suite.productStrepoMock.On("GetByIdProductSt", id).Return(id)
	value := ProductSt.FindProductStById(id)
	assert.Equal(suite.T(), value, dummyProductSt[0])
}
func (suite *ProductStUCTestSuite) TestFindIdProductSt_Failed() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	id := ""
	suite.productStrepoMock.On("GetByIdProductSt", id).Return(id)
	value := ProductSt.FindProductStById(id)
	assert.Equal(suite.T(), value, "productSt not found")
}

// ======== Register =========
func (suite *ProductStUCTestSuite) TestRegisterProductStAndSo_Success() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("CreateProductStAndSo", &dummyProductSt[0]).Return("register product to store and stock opname success")
	value := ProductSt.RegisterProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, "register product to store and stock opname success")
}
func (suite *ProductStUCTestSuite) TestRegisterProductStAndSo_Failed() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("CreateProductStAndSo", &dummyProductSt[0]).Return("register failed")
	value := ProductSt.RegisterProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, "register failed")
}
func (suite *ProductStUCTestSuite) TestRegisterProductStAndSo_Failed2() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("CreateProductStAndSo", &dummyProductSt[0]).Return("product already exist")
	value := ProductSt.RegisterProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, "product already exist")
}

// ======== Update =========
func (suite *ProductStUCTestSuite) TestUpdateProductSt_Success() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("UpdateProductSt", &dummyProductSt[0]).Return(fmt.Sprintf("ProductSt with id %s updated successfully", dummyProductSt[0].ID))
	value := ProductSt.EditProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, fmt.Sprintf("ProductSt with id %s updated successfully", dummyProductSt[0].ID))
}
func (suite *ProductStUCTestSuite) TestUpdateProductSt_Failed() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("UpdateProductSt", &dummyProductSt[0]).Return("product not found")
	value := ProductSt.EditProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, "product not found")
}
func (suite *ProductStUCTestSuite) TestUpdateProductSt_Failed2() {
	ProductSt := NewProductStUsecase(suite.productStrepoMock)
	suite.productStrepoMock.On("UpdateProductSt", &dummyProductSt[0]).Return("product not found")
	value := ProductSt.EditProductSt(&dummyProductSt[0])
	assert.Equal(suite.T(), value, "product not found")
}

//================================================================

// ============== setup first call=====================
func (suite *ProductStUCTestSuite) SetupTest() {
	suite.productStrepoMock = new(productStrepoMock)
}

func TestProductStUCTestSuite(t *testing.T) {
	suite.Run(t, new(ProductStUCTestSuite))
}
