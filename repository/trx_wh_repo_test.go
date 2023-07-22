package repository

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"go_inven_ctrl/entity"
)

var dummyTrx = []entity.TrxWh{
	{
		ID:          "TRXWH001",
		ProductWhId: "P001",
		ProductName: "Oreo",
		Stock:       1000,
		Act:         "input",
		CreatedAt:   time.Now(),
	},
	{
		ID:          "TRXWH002",
		ProductWhId: "P002",
		ProductName: "Beng-beng",
		Stock:       1000,
		Act:         "output",
		CreatedAt:   time.Now(),
	},
}

type TrxWhRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *TrxWhRepoTestSuite) TestEnrollInsertTrxWh_NotFound() {
	trxWh := dummyTrx[0]
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectCommit()

	repo := NewTrxWhRepo(suite.mockDb)

	actual := repo.EnrollInsertTrxWh(&trxWh)
	assert.Equal(suite.T(), "product not found", actual)
}

func (suite *TrxWhRepoTestSuite) TestEnrollInsertTrxWh_Rollback() {
	trxWh := dummyTrx[0]

	suite.mockSql.ExpectCommit()

	repo := NewTrxWhRepo(suite.mockDb)
	actual := repo.EnrollInsertTrxWh(&trxWh)

	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), "Transaction rollback", actual)
}

func (suite *TrxWhRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *TrxWhRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestTrxWhRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TrxWhRepoTestSuite))
}
