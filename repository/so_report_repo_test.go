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

// reportSo:= []entity.

var soReport = []entity.ReportSoRes{
	{ID: 1, TotalLoss: 120000, ProductMin: "Tango", TotalMin: 30000, ProductMax: "Frestea", TotalMax: 50000},
	{ID: 2, TotalLoss: 120000, ProductMin: "Oreo", TotalMin: 20000, ProductMax: "Frestea", TotalMax: 450000},
	{ID: 3, TotalLoss: 120000, ProductMin: "Taro", TotalMin: 25000, ProductMax: "Frestea", TotalMax: 40000},
}
var SoReportNIl = []entity.ReportSoRes{}

type SoReportRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *SoReportRepoTestSuite) TestGetAllReportInterim_Success() {
	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from interim_so_report order by id"
	rows := sqlmock.NewRows([]string{"id", "total_loss", "product_min", "total_min", "product_max", "total_max", "created_at"})
	for _, v := range soReport {
		rows.AddRow(v.ID, v.TotalLoss, v.ProductMin, v.TotalMin, v.ProductMax, v.TotalMax, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportSoRepo(suite.mockDb)
	actual := repo.GetAllInterimSoReport()
	assert.Equal(suite.T(), actual, soReport)
}

func (suite *SoReportRepoTestSuite) TestGetAllReportInterim_Failed() {
	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from interim_so_report order by id"
	rows := sqlmock.NewRows([]string{"id", "total_loss", "product_min", "total_min", "product_max", "total_max", "created_at"})
	for _, v := range SoReportNIl {
		rows.AddRow(v.ID, v.TotalLoss, v.ProductMin, v.TotalMin, v.ProductMax, v.TotalMax, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportSoRepo(suite.mockDb)
	actual := repo.GetAllInterimSoReport()
	assert.Equal(suite.T(), actual, "no data")
}

// ============================================================================================================
func (suite *SoReportRepoTestSuite) TestGetAllReportDetail_Success() {
	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from report_so_detail order by created_at"
	rows := sqlmock.NewRows([]string{"id", "total_loss", "product_min", "total_min", "product_max", "total_max", "created_at"})
	for _, v := range soReport {
		rows.AddRow(v.ID, v.TotalLoss, v.ProductMin, v.TotalMin, v.ProductMax, v.TotalMax, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportSoRepo(suite.mockDb)
	actual := repo.GetAlDetailSoReport()
	assert.Equal(suite.T(), actual, soReport)
}
func (suite *SoReportRepoTestSuite) TestGetAllReportDetail_Failed() {
	query := "select id,total_loss, product_min,total_min,product_max,total_max,created_at from report_so_detail order by created_at"
	rows := sqlmock.NewRows([]string{"id", "total_loss", "product_min", "total_min", "product_max", "total_max", "created_at"})
	for _, v := range SoReportNIl {
		rows.AddRow(v.ID, v.TotalLoss, v.ProductMin, v.TotalMin, v.ProductMax, v.TotalMax, v.CreatedAt)
	}
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewReportSoRepo(suite.mockDb)
	actual := repo.GetAlDetailSoReport()
	assert.Equal(suite.T(), actual, "no data")
}

//=======================================================================================================================

func (suite *SoReportRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func TestSoReportTestSuite(t *testing.T) {
	suite.Run(t, new(SoReportRepoTestSuite))
}
func (suite *SoReportRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
