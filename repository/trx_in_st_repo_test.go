package repository

import (
	"database/sql"
	"go_inven_ctrl/entity"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TrxStRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

var dummyTrxtSt = []entity.TrxInST{
	{
		ID: "1", ProductStId: "pst1", StockIn: 20, Act: "input", CreatedAt: time.Now(),
	},
	{
		ID: "2", ProductStId: "pst2", StockIn: 20, Act: "input", CreatedAt: time.Now(),
	},
	{
		ID: "3", ProductStId: "pst3", StockIn: 10, Act: "sold", CreatedAt: time.Now(),
	},
	{
		ID: "4", ProductStId: "pst4", StockIn: 10, Act: "retur", CreatedAt: time.Now(),
	},
}

func (suite *TrxStRepoTestSuite) TestEnrollTrxSt_Failed() {
	product := dummyTrxtSt[0]

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectCommit()

	productRepo := NewTrxInStRepo(suite.mockDb)
	str := productRepo.EnrollInsertTrxInSt(&product)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, "Transaction rollback")

}

func (suite *TrxStRepoTestSuite) TestEnrollTrxSt_Failed2() {
	product := dummyTrxtSt[0]

	suite.mockSql.ExpectCommit()

	productRepo := NewTrxInStRepo(suite.mockDb)
	str := productRepo.EnrollInsertTrxInSt(&product)
	assert.NotNil(suite.T(), str)
	assert.Equal(suite.T(), str, "product not found")

}

// ============== Before each =====================
func (suite *TrxStRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}
func TestTrxStTestSuite(t *testing.T) {
	suite.Run(t, new(TrxStRepoTestSuite))
}
func (suite *TrxStRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
