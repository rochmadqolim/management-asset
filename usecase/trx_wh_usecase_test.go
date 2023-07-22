package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyTrx = []entity.TrxWh{
	{
		ID:          "TRXWH001",
		ProductWhId: "P001",
		ProductName: "Oreo",
		Stock:       1000,
		Act:         "input",
	},
	{
		ID:          "TRXWH002",
		ProductWhId: "P002",
		ProductName: "Beng-beng",
		Stock:       1000,
		Act:         "output",
	},
}

type repoMockTrxWh struct {
	mock.Mock
}

type TrxWhUsecaseTestSuite struct {
	repoMockTrxWh *repoMockTrxWh
	suite.Suite
}

func (r *repoMockTrxWh) EnrollInsertTrxWh(enrollTrxWh *entity.TrxWh) string {
	args := r.Called(enrollTrxWh)
	if args.Get(0) != nil {
		return args.String(0)
	}
	return "transaction in warehouse committed"
}

func (suite *TrxWhUsecaseTestSuite) TestEnrollInsertTrxWh_Success() {
	trxWhUc := NewTrxWhUsecase(suite.repoMockTrxWh)

	suite.repoMockTrxWh.On("EnrollInsertTrxWh", &dummyTrx[0]).Return("transaction in warehouse committed")

	trxWh := trxWhUc.EnrollInsertTrxWh(&dummyTrx[0])

	assert.Equal(suite.T(), "transaction in warehouse committed", trxWh)
}

func (suite *TrxWhUsecaseTestSuite) TestEnrollInsertTrxWh_Failed() {
	trxWhUc := NewTrxWhUsecase(suite.repoMockTrxWh)

	suite.repoMockTrxWh.On("EnrollInsertTrxWh", &dummyTrx[0]).Return("Transaction rollback")

	trxWh := trxWhUc.EnrollInsertTrxWh(&dummyTrx[0])

	assert.Equal(suite.T(), "Transaction rollback", trxWh)
}

func (suite *TrxWhUsecaseTestSuite) SetupTest() {
	suite.repoMockTrxWh = new(repoMockTrxWh)
}

func TestTrxWhUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TrxWhUsecaseTestSuite))
}
