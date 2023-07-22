package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyProductSo = []entity.ProductSo{
	{ID: "so1", ProductStId: "pst1", Stock: 50, DiffStock: -10, DiffPrice: 50000},
	{ID: "so2", ProductStId: "pst2", Stock: 45, DiffStock: -5, DiffPrice: 20000},
	{ID: "so3", ProductStId: "pst3", Stock: 40, DiffStock: -10, DiffPrice: 50000},
}

type productSoRepoMock struct {
	mock.Mock
}

func (r *productSoRepoMock) GetAllProductSo() any {
	args := r.Called()
	if args[0] == nil {
		return "no data"
	}
	return dummyProductSo
}
func (r *productSoRepoMock) GetByLessThan(stock int) any {
	args := r.Called(stock)
	if args[0] == nil {
		return "no data"
	}
	return dummyProductSo
}

type ProductSoUCTestSuite struct {
	productSoRepoMock *productSoRepoMock
	suite.Suite
}

// ====================================Testing=====================================================
// GetAll
func (suite *ProductSoUCTestSuite) TestFindAllProductSo_Success() {
	productSo := NewProductSoUsecase(suite.productSoRepoMock)
	suite.productSoRepoMock.On("GetAllProductSo").Return(dummyProductSo)
	value := productSo.FindAllProductSo()
	assert.Equal(suite.T(), value, dummyProductSo)
}

func (suite *ProductSoUCTestSuite) TestFindAllProductSo_success() {
	productSo := NewProductSoUsecase(suite.productSoRepoMock)
	suite.productSoRepoMock.On("GetAllProductSo").Return(nil)
	value := productSo.FindAllProductSo()
	assert.Equal(suite.T(), value, "no data")
}

// GetByLessThan
func (suite *ProductSoUCTestSuite) TestFindByLessThanProductSo_Success() {
	var stock = -5
	productSo := NewProductSoUsecase(suite.productSoRepoMock)
	suite.productSoRepoMock.On("GetByLessThan", stock).Return(dummyProductSo)
	value := productSo.FindByLessThan(stock)
	assert.Equal(suite.T(), value, dummyProductSo)
}

func (suite *ProductSoUCTestSuite) TestFindByLessThanProductSo_success() {
	var stock = -10
	productSo := NewProductSoUsecase(suite.productSoRepoMock)
	suite.productSoRepoMock.On("GetByLessThan", stock).Return(nil)
	value := productSo.FindByLessThan(stock)
	assert.Equal(suite.T(), value, "no data")
}

// ==================================Setup ===================================================
func (suite *ProductSoUCTestSuite) SetupTest() {
	suite.productSoRepoMock = new(productSoRepoMock)
}

func TestProductSoUCTestSuite(t *testing.T) {
	suite.Run(t, new(ProductSoUCTestSuite))
}
