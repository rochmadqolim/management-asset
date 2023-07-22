package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type trxProductSoRepoMock struct {
	mock.Mock
}

var dummyTrxSo = []entity.TrxInSo{
	{ID: 1, ProductStSoId: "pst1", Stock: 48},
	{ID: 1, ProductStSoId: "pst2", Stock: 38},
	{ID: 1, ProductStSoId: "pst3", Stock: 40},
}

func (r *trxProductSoRepoMock) EnrollInsertTrxInSo(EnrollTrxInSt *entity.TrxInSo) string {
	args := r.Called(EnrollTrxInSt)
	if args[0] == nil {
		return "product not found"
	}
	if args[0] == "Transaction rollback" {
		return "Transaction rollback"
	}
	return "Transaction in stock opname  commited"
}
func (r *trxProductSoRepoMock) EnrollInsertReportConfirm() string {
	args := r.Called()
	if args[0] == nil {
		return "product not found"
	}
	if args[0] == "Transaction rollback" {
		return "Transaction rollback"
	}
	return "Transaction in stock opname  commited"
}
func (r *trxProductSoRepoMock) EnrollInsertReportInterim() string {
	args := r.Called()
	if args[0] == nil {
		return "product not found"
	}
	if args[0] == "Transaction rollback" {
		return "Transaction rollback"
	}

	return "Transaction in stock opname  commited"
}

type TrxProductSoUCTestSuite struct {
	trxProductSoRepoMock *trxProductSoRepoMock
	suite.Suite
}

func (suite *TrxProductSoUCTestSuite) TestEnrollInsertTrxSo_Success() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertTrxInSo", &dummyTrxSo[0]).Return(&dummyTrxSo[0])
	value := trxProductSo.EnrollInsertTrxSo(&dummyTrxSo[0])
	assert.Equal(suite.T(), value, "Transaction in stock opname  commited")
}
func (suite *TrxProductSoUCTestSuite) TestEnrollInsertTrxSo_Failed() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertTrxInSo", &dummyTrxSo[0]).Return(nil)
	value := trxProductSo.EnrollInsertTrxSo(&dummyTrxSo[0])
	assert.Equal(suite.T(), value, "product not found")
}
func (suite *TrxProductSoUCTestSuite) TestEnrollInsertTrxSo_Failed2() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertTrxInSo", &dummyTrxSo[0]).Return("Transaction rollback")
	value := trxProductSo.EnrollInsertTrxSo(&dummyTrxSo[0])
	assert.Equal(suite.T(), value, "Transaction rollback")
}

func (suite *TrxProductSoUCTestSuite) TestReportConfirmation_Success() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportConfirm").Return(&dummyTrxSo[0])
	value := trxProductSo.ReportConfirmation()
	assert.Equal(suite.T(), value, "Transaction in stock opname  commited")
}
func (suite *TrxProductSoUCTestSuite) TestReportConfirmation_Failed() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportConfirm").Return(nil)
	value := trxProductSo.ReportConfirmation()
	assert.Equal(suite.T(), value, "product not found")
}
func (suite *TrxProductSoUCTestSuite) TestReportConfirmation_Failed2() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportConfirm").Return("Transaction rollback")
	value := trxProductSo.ReportConfirmation()
	assert.Equal(suite.T(), value, "Transaction rollback")
}

func (suite *TrxProductSoUCTestSuite) TestReportInterim_Success() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportInterim").Return(&dummyTrxSo[0])
	value := trxProductSo.ReportInterim()
	assert.Equal(suite.T(), value, "Transaction in stock opname  commited")
}
func (suite *TrxProductSoUCTestSuite) TestReportInterim_Failed() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportInterim").Return(nil)
	value := trxProductSo.ReportInterim()
	assert.Equal(suite.T(), value, "product not found")
}
func (suite *TrxProductSoUCTestSuite) TestReportInterim_Failed2() {
	trxProductSo := NewTrxInSoUsecase(suite.trxProductSoRepoMock)
	suite.trxProductSoRepoMock.On("EnrollInsertReportInterim").Return("Transaction rollback")
	value := trxProductSo.ReportInterim()
	assert.Equal(suite.T(), value, "Transaction rollback")
}

// ==========================================================================================
func (suite *TrxProductSoUCTestSuite) SetupTest() {
	suite.trxProductSoRepoMock = new(trxProductSoRepoMock)
}

func TestTrxProductSoUCTestSuite(t *testing.T) {
	suite.Run(t, new(TrxProductSoUCTestSuite))
}
