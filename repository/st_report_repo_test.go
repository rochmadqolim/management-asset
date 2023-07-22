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

type ReportStRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

var reportTrxSt = []entity.ReportTrxSt{
	{
		ProductStId: "pst1", StockIn: 50, ProductName: "Ultra Milk", Act: "input", LastStock: 0, CreatedAt: "2023-04-10",
	},
	{
		ProductStId: "pst1", StockIn: 30, ProductName: "Ultra Milk", Act: "sold", LastStock: 0, CreatedAt: "2023-04-10",
	},
	{
		ProductStId: "pst1", StockIn: 20, ProductName: "Ultra Milk", Act: "retur", LastStock: 0, CreatedAt: "2023-04-10",
	},
}
var reportTrxStNil = []entity.ReportTrxSt{}

// ====================================find By Id ==================================================
func (suite *ReportStRepoTestSuite) TestGetByProductStId_Success() {

	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where product_st_id = \\$1 order by act asc, created_at asc"
	id := "pst1"
	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxSt {
		if v.ProductStId == id {
			rows.AddRow(v.ProductStId, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetByReportTrxProductStId(id)
	assert.Equal(suite.T(), actual, reportTrxSt)
}

func (suite *ReportStRepoTestSuite) TestGetByProductStId_Failed2() {

	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where product_st_id = \\$1 order by act asc, created_at asc"
	id := "pst4"
	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxSt {
		if v.ProductStId == id {
			rows.AddRow(v.ProductStId, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetByReportTrxProductStId(id)
	assert.Equal(suite.T(), actual, "no data or product id not found")

}

// =====================================find by date==============================================
func (suite *ReportStRepoTestSuite) TestGetByProductStDate_Success() {
	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where created_at = \\$1 order by act asc, created_at asc"

	date := "2023-04-10"
	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxSt {
		if v.CreatedAt == date {
			rows.AddRow(v.ProductStId, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(date).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetByDateReportTrxSt(date)
	assert.Equal(suite.T(), actual, reportTrxSt)
}

func (suite *ReportStRepoTestSuite) TestGetByProductStDate_Failed2() {
	query := "select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st where created_at = \\$1 order by act asc, created_at asc"

	date := "2023-04-12"
	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxSt {
		if v.CreatedAt == date {
			rows.AddRow(nil, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}
	suite.mockSql.ExpectQuery(query).WithArgs(date).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetByDateReportTrxSt(date)
	assert.Equal(suite.T(), actual, "no data or date product not found")
	log.Println(actual)

}

// =====================================Get All==============================================
func (suite *ReportStRepoTestSuite) TestGetAll_Success() {
	query := `select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st order by act asc, created_at asc`

	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxSt {
		rows.AddRow(v.ProductStId, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetAllReportTrxSt()
	assert.Equal(suite.T(), actual, reportTrxSt)
}
func (suite *ReportStRepoTestSuite) TestGetAll_Failed() {
	query := `select  product_st_id, stock_in, product_name, act, last_stock,created_at from report_trx_st order by act asc, created_at asc`

	rows := sqlmock.NewRows([]string{"product_st_id", "stock_in", "product_name", "act", "last_stock", "created_at"})
	for _, v := range reportTrxStNil {
		rows.AddRow(nil, v.StockIn, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportTrxStRepo(suite.mockDb)
	actual := repo.GetAllReportTrxSt()
	assert.Equal(suite.T(), actual, "no data")
}

// ============== Before each =====================
func (suite *ReportStRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}
func TestReportStTestSuite(t *testing.T) {
	suite.Run(t, new(ReportStRepoTestSuite))
}
func (suite *ReportStRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
