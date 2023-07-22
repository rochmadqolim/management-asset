package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go_inven_ctrl/entity"

	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyProducts = []entity.ProductWh{
	{
		ID:              "1",
		ProductName:     "Oreo",
		Price:           3000,
		ProductCategory: "food",
		Stock:           1000,
	},
	{
		ID:              "2",
		ProductName:     "Beng-beng",
		Price:           2500,
		ProductCategory: "food",
		Stock:           1000,
	},
	{
		ID:              "3",
		ProductName:     "Indomie",
		Price:           3500,
		ProductCategory: "food",
		Stock:           1000,
	},
}

type ProductWhRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *ProductWhRepoTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})

	for _, v := range dummyProducts {
		rows.AddRow(v.ID, v.ProductName, v.Price, v.ProductCategory, v.Stock)
	}

	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh").WillReturnRows(rows)

	repo := NewProductWhRepo(suite.mockDb)
	expected := dummyProducts
	actual := repo.GetAll().([]entity.ProductWh)

	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), len(dummyProducts), len(actual))
	assert.Equal(suite.T(), "2", actual[1].ID)
}

func (suite *ProductWhRepoTestSuite) TestGetAllScan_Failed() {
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})
	rows.AddRow(nil, "Yakult", 2500, "food", 1000)

	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh").WillReturnRows(rows)

	repo := NewProductWhRepo(suite.mockDb)
	actual := repo.GetAll()

	expected := []entity.ProductWh{
		{ID: "", ProductName: "", Price: 0, ProductCategory: "", Stock: 0},
	}
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestGetAll_Empty() {
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"})

	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh").WillReturnRows(rows)

	repo := NewProductWhRepo(suite.mockDb)
	actual := repo.GetAll()

	assert.Equal(suite.T(), "no data", actual)
}

func (suite *ProductWhRepoTestSuite) TestGetById_Success() {
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(dummyProducts[0].ID, dummyProducts[0].ProductName, dummyProducts[0].Price, dummyProducts[0].ProductCategory, dummyProducts[0].Stock)

	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs("1").WillReturnRows(rows)

	repo := NewProductWhRepo(suite.mockDb)
	expected := dummyProducts[0]
	actual := repo.GetById("1")

	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestGetById_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs("product not found").WillReturnError(errors.New("Failed to get product"))

	repo := NewProductWhRepo(suite.mockDb)
	actual := repo.GetById("invalid-id")

	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), "product not found", actual)
}

func (suite *ProductWhRepoTestSuite) TestCreateProductWh_Success() {
	newProduct := dummyProducts[0]
	suite.mockSql.ExpectExec("INSERT INTO product_wh\\(id, product_name, price, product_category, stock\\) VALUES").WithArgs(
		newProduct.ID,
		newProduct.ProductName,
		newProduct.Price,
		newProduct.ProductCategory,
		newProduct.Stock,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductWhRepo(suite.mockDb)
	actual := repo.Create(&newProduct)

	assert.Equal(suite.T(), "new product created successfully", actual)
}

func (suite *ProductWhRepoTestSuite) TestCreateProductWh_Failed() {
	newProduct := dummyProducts[0]
	suite.mockSql.ExpectExec("INSERT INTO product_wh\\(id, product_name, price, product_category, stock\\) VALUES").WillReturnError(errors.New("failed to create product"))

	repo := NewProductWhRepo(suite.mockDb)
	actual := repo.Create(&newProduct)

	assert.Equal(suite.T(), "failed to create product", actual)
}

func (suite *ProductWhRepoTestSuite) TestProductWhUpdate_Success() {
	rows := sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(dummyProducts[0].ID, dummyProducts[0].ProductName, dummyProducts[0].Price, dummyProducts[0].ProductCategory, dummyProducts[0].Stock)

	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ID).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE product_wh SET id = \\$1, product_name = \\$2, price = \\$3, product_category = \\$4, stock = \\$5 WHERE id = \\$6").WithArgs(dummyProducts[0].ID, dummyProducts[0].ProductName, 8000, dummyProducts[0].ProductCategory, dummyProducts[0].Stock, dummyProducts[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewProductWhRepo(suite.mockDb)
	updatedProduct := &entity.ProductWh{
		ID:              dummyProducts[0].ID,
		ProductName:     dummyProducts[0].ProductName,
		Price:           8000,
		ProductCategory: dummyProducts[0].ProductCategory,
		Stock:           dummyProducts[0].Stock,
	}
	expected := fmt.Sprintf("product with ID %s updated successfully", updatedProduct.ID)
	actual := repo.Update(updatedProduct)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestProductWhUpdate_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ProductName).WillReturnError(errors.New("product not found"))

	repo := NewProductWhRepo(suite.mockDb)
	expected := "product not found"
	actual := repo.Update(&dummyProducts[0])

	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestProductWhDelete_Success() {
	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}).AddRow(dummyProducts[0].ID, "Oreo", 3000, "food", 350))

	suite.mockSql.ExpectExec("DELETE FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewProductWhRepo(suite.mockDb)
	expected := fmt.Sprintf("product with id %s deleted successfully", dummyProducts[0].ID)
	actual := repo.Delete(dummyProducts[0].ID)
	if actual != expected {
		err := repo.GetById(dummyProducts[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestProductWhDelete_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "price", "product_category", "stock"}))

	repo := NewProductWhRepo(suite.mockDb)
	expected := "product not found"
	actual := repo.Delete(dummyProducts[0].ID)

	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) TestProductWhDelete_NotFound() {
	suite.mockSql.ExpectQuery("SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = ?").WithArgs(dummyProducts[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewProductWhRepo(suite.mockDb)
	expected := "product not found"
	actual := repo.Delete(dummyProducts[0].ID)

	assert.Equal(suite.T(), expected, actual)
}

func (suite *ProductWhRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *ProductWhRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestProductWhRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductWhRepoTestSuite))
}
