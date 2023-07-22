package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type trxProductStRepoMock struct {
	mock.Mock
}

var dummyTrxSt = []entity.TrxInST{
	{ProductStId: "pst1", StockIn: 50, Act: "input"},
	{ProductStId: "pst1", StockIn: 10, Act: "sold"},
	{ProductStId: "pst1", StockIn: 2, Act: "retur"},
}

func (r *trxProductStRepoMock) EnrollInsertTrxInSt(EnrollTrxInSt *entity.TrxInST) string {
	args := r.Called(EnrollTrxInSt)
	if args[0] == nil {
		return "product not found"
	}
	if args[0] == "Transaction rollback" {
		return "Transaction rollback"
	}
	return "Transaction in store  commited"
}

type TrxProductStUCTestSuite struct {
	trxProductStRepoMock *trxProductStRepoMock
	suite.Suite
}

func (suite *TrxProductStUCTestSuite) TestTrxProductSt_Success() {
	trxProductSt := NewTrxInStUsecase(suite.trxProductStRepoMock)
	suite.trxProductStRepoMock.On("EnrollInsertTrxInSt", &dummyTrxSt[0]).Return(&dummyTrxSt[0])
	value := trxProductSt.EnrollInsertTrx(&dummyTrxSt[0])
	assert.Equal(suite.T(), value, "Transaction in store  commited")
}
func (suite *TrxProductStUCTestSuite) TestTrxProductSt_Failed() {
	trxProductSt := NewTrxInStUsecase(suite.trxProductStRepoMock)
	suite.trxProductStRepoMock.On("EnrollInsertTrxInSt", &dummyTrxSt[0]).Return(nil)
	value := trxProductSt.EnrollInsertTrx(&dummyTrxSt[0])
	assert.Equal(suite.T(), value, "product not found")
}
func (suite *TrxProductStUCTestSuite) TestTrxProductSt_Failed2() {
	trxProductSt := NewTrxInStUsecase(suite.trxProductStRepoMock)
	suite.trxProductStRepoMock.On("EnrollInsertTrxInSt", &dummyTrxSt[0]).Return("Transaction rollback")
	value := trxProductSt.EnrollInsertTrx(&dummyTrxSt[0])
	assert.Equal(suite.T(), value, "Transaction rollback")
}

func (suite *TrxProductStUCTestSuite) SetupTest() {
	suite.trxProductStRepoMock = new(trxProductStRepoMock)
}

func TestTrxProductStUCTestSuite(t *testing.T) {
	suite.Run(t, new(TrxProductStUCTestSuite))
}
