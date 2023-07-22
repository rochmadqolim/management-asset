package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductStRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

var dummyProductSt = []entity.ProductSt{
	{
		ID: "pst1", ProductName: "Oreo", ProductCtg: "makanan", Price: 5000, Stock: 0,
	},
	{
		ID: "pst2", ProductName: "Richeese", ProductCtg: "makanan", Price: 3000, Stock: 0,
	},
	{
		ID: "pst3", ProductName: "Teh Kotak", ProductCtg: "minuman", Price: 4000, Stock: 0,
	},
}
var dummynil = []entity.ProductSt{}

func (suite *ProductStRepoTestSuite) TestUpdateProductSt_Success() {
	product := dummyProductSt[0]
	queryUpdate := "UPDATE product_st SET product_name = \\$1, price = \\$2,product_category = \\$3 WHERE id = \\$4 ;"

	suite.mockSql.ExpectQuery("SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = \\$1").
		WithArgs(product.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(
		product.ID, product.ProductName, product.Price, product.ProductCtg, product.Stock,
	))
	suite.mockSql.ExpectExec(queryUpdate).WithArgs(
		product.ProductName, product.Price, product.ProductCtg, product.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	productRepo := NewProductStRepo(suite.mockDb)
	str := productRepo.UpdateProductSt(&product)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, fmt.Sprintf("ProductSt with id %s updated successfully", product.ID))

}
func (suite *ProductStRepoTestSuite) TestUpdateProductSt_Failed() {
	product := dummyProductSt[2]

	queryUpdate := "UPDATE product_st SET product_name = \\$1, price = \\$2,product_category = \\$3 WHERE id = \\$4 ;"
	suite.mockSql.ExpectExec(queryUpdate).WithArgs(
		product.ProductName, product.Price, product.ProductCtg, product.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	productRepo := NewProductStRepo(suite.mockDb)
	str := productRepo.UpdateProductSt(&product)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, "productSt not found")

}
func (suite *ProductStRepoTestSuite) TestUpdateProductSt_Failed2() {
	product := dummyProductSt[2]

	suite.mockSql.ExpectQuery("SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = \\$1").
		WithArgs(product.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product", "price", "product_category", "stock"}).AddRow(
		product.ID, product.ProductName, product.Price, product.ProductCtg, product.Stock,
	))

	productRepo := NewProductStRepo(suite.mockDb)
	str := productRepo.UpdateProductSt(&product)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, "failed to update ProductSt")
	log.Println(str)

}

// =================================Get All ProductSt=====================================================
func (suite *ProductStRepoTestSuite) TestGetAllProductSt_Success() {
	query := "SELECT  id,product_name,price,product_category,stock FROM product_st order by id asc"
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	for _, v := range dummyProductSt {
		rows.AddRow(v.ID, v.ProductName, v.Price, v.ProductCtg, v.Stock)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewProductStRepo(suite.mockDb)
	actual := repo.GetAllProductSt()
	assert.Equal(suite.T(), actual, dummyProductSt)
}
func (suite *ProductStRepoTestSuite) TestGetAllProductSt_Failed() {
	query := "SELECT  id,product_name,price,product_category,stock FROM product_st order by id asc"
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	for _, v := range dummynil {
		rows.AddRow(v.ID, v.ProductName, v.Price, v.ProductCtg, nil)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewProductStRepo(suite.mockDb)
	actual := repo.GetAllProductSt()
	assert.Equal(suite.T(), actual, "no data")
	log.Println(actual)
}

func (suite *ProductStRepoTestSuite) TestGetAllProductSt_Failed2() {
	query := "SELECT  id,product_name,price,product_category,stock FROM product_st order by id asc"
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	rows.AddRow(nil, dummyProductSt[0].ProductName, dummyProductSt[0].Price, dummyProductSt[0].ProductCtg, dummyProductSt[0].Stock)
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewProductStRepo(suite.mockDb)
	actual := repo.GetAllProductSt()
	log.Println(actual)
}

// =================================== Get By Id =======================================

func (suite *ProductStRepoTestSuite) TestGetById_Success() {
	query := "SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id ="
	id := "pst3"
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	for _, v := range dummyProductSt {
		if v.ID == id {
			rows.AddRow(v.ID, v.ProductName, v.Price, v.ProductCtg, v.Stock)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	repo := NewProductStRepo(suite.mockDb)
	actual := repo.GetByIdProductSt(id)
	assert.Equal(suite.T(), actual, dummyProductSt[2])
}
func (suite *ProductStRepoTestSuite) TestGetById_Failed() {
	query := "SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id ="
	id := "pst4"
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	for _, v := range dummyProductSt {
		if v.ID == id {
			rows.AddRow(v.ID, v.ProductName, v.Price, v.ProductCtg, v.Stock)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	repo := NewProductStRepo(suite.mockDb)
	actual := repo.GetByIdProductSt(id)
	assert.Equal(suite.T(), actual, "productSt not found")
}

// =================================Delete ProductSt and So=====================================================
func (suite *ProductStRepoTestSuite) TestDeleteProductStAndSo_Success() {
	product := dummyProductSt[0]
	id := "pst1"
	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectQuery("SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = \\$1").
		WithArgs(product.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(
		product.ID, product.ProductName, product.Price, product.ProductCtg, product.Stock,
	))

	suite.mockSql.ExpectExec("DELETE FROM product_so WHERE product_st_id =\\$1").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mockSql.ExpectQuery("SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = \\$1").
		WithArgs(product.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(
		product.ID, product.ProductName, product.Price, product.ProductCtg, product.Stock,
	))

	suite.mockSql.ExpectExec("DELETE FROM product_st WHERE id =\\$1").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()

	productRepo := NewProductStRepo(suite.mockDb)
	str := productRepo.DeleteProductStAndSo(id)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, fmt.Sprintf("delete product with id %s from store and stock opname success", id))

}

func (suite *ProductStRepoTestSuite) TestDeleteProductStAndSo_Failed() {
	query := "DELETE FROM product_so WHERE product_st_id =$1"
	query2 := "DELETE FROM product_st WHERE id =$1"
	newProductSt := dummyProductSt[0]
	//==============transaction====================
	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectExec(query).WithArgs(
		newProductSt.ID,
	)
	suite.mockSql.ExpectExec(query2).WithArgs(
		newProductSt.ID,
	)

	suite.mockSql.ExpectCommit()

	productStRepo := NewProductStRepo(suite.mockDb)
	value := productStRepo.DeleteProductStAndSo(newProductSt.ID)
	assert.Equal(suite.T(), value, "Failed to delete product")
	log.Println(value)
}
func (suite *ProductStRepoTestSuite) TestDeleteProductStAndSo_Failed2() {
	query := "DELETE FROM product_st WHERE id =$1"
	newProductSt := dummyProductSt[2]
	//==============transaction====================

	suite.mockSql.ExpectExec(query).WithArgs(
		newProductSt.ID,
	)

	suite.mockSql.ExpectCommit()

	productStRepo := NewProductStRepo(suite.mockDb)
	value := productStRepo.DeleteProductStAndSo(newProductSt.ID)
	log.Println(value)
	assert.Equal(suite.T(), value, "product not found")
}

// =================================Create ProductSt and So=====================================================
func (suite *ProductStRepoTestSuite) TestCreateProductStAndSo_Success() {
	newProductSt := &dummyProductSt[0]
	query := "INSERT INTO product_st \\( id, product_name, product_category, price\\) VALUES"
	query2 := "INSERT INTO product_so \\(product_st_id\\) VALUES"
	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectExec(query).WithArgs(newProductSt.ID, newProductSt.ProductName, newProductSt.ProductCtg, newProductSt.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectExec(query2).WithArgs(newProductSt.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mockSql.ExpectCommit()
	productStRepo := NewProductStRepo(suite.mockDb)
	value := productStRepo.CreateProductStAndSo(newProductSt)
	fmt.Println(value)
	assert.Equal(suite.T(), value, "register product to store and stock opname success")
}

func (suite *ProductStRepoTestSuite) TestCreateProductStAndSo_Failed() {
	newProductSt := dummyProductSt[0]
	suite.mockSql.ExpectExec("INSERT INTO product_st ( id, product_name, product_category, price) VALUES").WithArgs(
		newProductSt.ID, newProductSt.ProductName,
		newProductSt.ProductCtg, newProductSt.Price,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	productStRepo := NewProductStRepo(suite.mockDb)
	value := productStRepo.CreateProductStAndSo(&dummyProductSt[0])
	fmt.Println(value)
	assert.Equal(suite.T(), value, "product already exist")
}
func (suite *ProductStRepoTestSuite) TestCreateProductStAndSo_Failed2() {
	newProductSt := dummyProductSt[0]
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO product_st ( id, product_name, product_category, price) VALUES").WithArgs(
		newProductSt.ID, newProductSt.ProductName,
		newProductSt.ProductCtg, newProductSt.Price,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit()
	productStRepo := NewProductStRepo(suite.mockDb)
	value := productStRepo.CreateProductStAndSo(&dummyProductSt[0])
	fmt.Println(value)
	assert.Equal(suite.T(), value, "register failed")
}

// ============== Before each =====================
func (suite *ProductStRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}
func TestProductStRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductStRepoTestSuite))
}
func (suite *ProductStRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
