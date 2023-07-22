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

var dummyReportWh = []entity.ReportWh{
	{
		ProductWhId: "P001",
		Stock:       100,
		ProductName: "Oreo",
		Act:         "input",
		LastStock:   1000,
		CreatedAt:   "2023-01-01",
	},
}

type WhReportRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *WhReportRepoTestSuite) TestGetAllReportTrxWh_Success() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})

	for _, v := range dummyReportWh {
		rows.AddRow(v.ProductWhId, v.Stock, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
	}

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh ORDER BY act ASC, created_at ASC").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	expected := dummyReportWh
	actual := repo.GetAllReportTrxWh().([]entity.ReportWh)

	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), len(dummyReportWh), len(actual))
	assert.Equal(suite.T(), "P002", actual[1].ProductWhId)
}

func (suite *WhReportRepoTestSuite) TestGetAllReportTrxWhScan_Failed() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})
	rows.AddRow(nil, 300, "Oreo", "input", 900, "2023-01-01")

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh ORDER BY act ASC, created_at ASC").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	actual := repo.GetAllReportTrxWh()

	expected := []entity.ReportWh{
		{ProductWhId: "", Stock: 0, ProductName: "", Act: "", LastStock: 0, CreatedAt: ""},
	}
	assert.Equal(suite.T(), expected, actual)
}

func (suite *WhReportRepoTestSuite) TestGetAllReportTrxWh_Empty() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh ORDER BY act ASC, created_at ASC").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	actual := repo.GetAllReportTrxWh()

	assert.Equal(suite.T(), "no data", actual)
}

func (suite *WhReportRepoTestSuite) TestGetByIdReportTrxWh_Success() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})
	for _, v := range dummyReportWh {
		if v.ProductWhId == "P001" {
			rows.AddRow(v.ProductWhId, v.Stock, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE product_wh_id = \\$1 ORDER BY act ASC, created_at ASC").WithArgs("P001").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	//expected := dummyReportWh[0]
	actual := repo.GetByIdReportTrxWh("P001")

	assert.Equal(suite.T(), dummyReportWh, actual)
}

func (suite *WhReportRepoTestSuite) TestGetByIdReportTrxWh_Failed() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})
	for _, v := range dummyReportWh {
		if v.ProductWhId == "P005" {
			rows.AddRow(v.ProductWhId, v.Stock, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE product_wh_id = \\$1 ORDER BY act ASC, created_at ASC").WithArgs("P005").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	actual := repo.GetByIdReportTrxWh("P005")

	assert.Equal(suite.T(), "no data or product id not found", actual)

}

func (suite *WhReportRepoTestSuite) TestGetByDateReportTrxWh_Success() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})
	for _, v := range dummyReportWh {
		if v.CreatedAt == "2023-01-01" {
			rows.AddRow(v.ProductWhId, v.Stock, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE created_at = \\$1 ORDER BY act ASC, created_at ASC").WithArgs("2023-01-01").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	actual := repo.GetByDateReportTrxWh("2023-01-01")

	assert.Equal(suite.T(), dummyReportWh, actual)
}

func (suite *WhReportRepoTestSuite) TestGetByDateReportTrxWh_Failed() {
	rows := sqlmock.NewRows([]string{"product_wh_id", "stock", "product_name", "act", "last_stock", "created_at"})
	for _, v := range dummyReportWh {
		if v.CreatedAt == "2000-01-01" {
			rows.AddRow(v.ProductWhId, v.Stock, v.ProductName, v.Act, v.LastStock, v.CreatedAt)
		}
	}

	suite.mockSql.ExpectQuery("SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE created_at = \\$1 ORDER BY act ASC, created_at ASC").WithArgs("2000-01-01").WillReturnRows(rows)

	repo := NewReportTrxWhRepo(suite.mockDb)
	actual := repo.GetByDateReportTrxWh("2000-01-01")

	assert.Equal(suite.T(), "no data or date product not found", actual)
}

func (suite *WhReportRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *WhReportRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestWhReportRepoTestSuite(t *testing.T) {
	suite.Run(t, new(WhReportRepoTestSuite))
}
