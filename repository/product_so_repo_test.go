package repository

import (
	"database/sql"
	"go_inven_ctrl/entity"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductSoRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

var reportTrxSo = []entity.ProductSo{
	{ID: "1", ProductStId: "pst1", Stock: 41, DiffStock: -9, DiffPrice: -45000},
	{ID: "1", ProductStId: "pst1", Stock: 40, DiffStock: -10, DiffPrice: -50000},
	{ID: "1", ProductStId: "pst1", Stock: 45, DiffStock: -5, DiffPrice: -25000},
}

var reportTrxSoNIl = []entity.ProductSo{}

//========================================================================Get All Product So====================================================

func (suite *ProductSoRepoTestSuite) TestGetAllProductSo_Success() {
	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where diff_stock \\<\\> 0 order by diff_price asc"

	rows := sqlmock.NewRows([]string{"pst.id", "pst.product_name", "pso.stock", "pso.diff_stock", "pso.diff_price"})
	for _, v := range reportTrxSo {

		rows.AddRow(v.ID, v.ProductStId, v.Stock, v.DiffStock, v.DiffPrice)

	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewProductSoRepo(suite.mockDb)
	actual := repo.GetAllProductSo()
	assert.Equal(suite.T(), actual, reportTrxSo)
}

func (suite *ProductSoRepoTestSuite) TestGetAllProductSo_Failed() {
	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where diff_stock \\<\\> 0 order by diff_price asc"

	rows := sqlmock.NewRows([]string{"pst.id", "pst.product_name", "pso.stock", "pso.diff_stock", "pso.diff_price"})
	for _, v := range reportTrxSoNIl {

		rows.AddRow(v.ID, v.ProductStId, v.Stock, v.DiffStock, v.DiffPrice)

	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewProductSoRepo(suite.mockDb)
	actual := repo.GetAllProductSo()
	assert.Equal(suite.T(), actual, "no data")
}

// ======================================================================= GEt Less Than=================================================
func (suite *ProductSoRepoTestSuite) TestGetLessThenProductSo_Success() {
	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where pso.diff_stock \\< \\$1 order by diff_price asc"

	diffStock := -8
	rows := sqlmock.NewRows([]string{"pst.id", "pst.product_name", "pso.stock", "pso.diff_price", "pso.diff_stock"})
	for _, v := range reportTrxSo {
		if v.DiffStock <= diffStock {
			rows.AddRow(v.ID, v.ProductStId, v.Stock, v.DiffStock, v.DiffPrice)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(diffStock).WillReturnRows(rows)
	repo := NewProductSoRepo(suite.mockDb)
	actual := repo.GetByLessThan(diffStock)
	assert.Equal(suite.T(), actual, reportTrxSo)
}
func (suite *ProductSoRepoTestSuite) TestGetLessThenProductSo_Failed() {
	query := "select pst.id,pst.product_name,pso.stock,pso.diff_price, pso.diff_stock from product_so as pso join product_st as pst on pso.product_st_id =pst.id where pso.diff_stock \\< \\$1 order by diff_price asc"

	diffStock := -10
	rows := sqlmock.NewRows([]string{"pst.id", "pst.product_name", "pso.stock", "pso.diff_price", "pso.diff_stock"})
	for _, v := range reportTrxSo {
		if v.DiffStock < diffStock {
			rows.AddRow(v.ID, v.ProductStId, v.Stock, v.DiffStock, v.DiffPrice)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(diffStock).WillReturnRows(rows)
	repo := NewProductSoRepo(suite.mockDb)
	actual := repo.GetByLessThan(diffStock)
	assert.Equal(suite.T(), actual, "no data")
}

//============================================================================================

func (suite *ProductSoRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}
func TestReportSotRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductSoRepoTestSuite))
}
func (suite *ProductSoRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
