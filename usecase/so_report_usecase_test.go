package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyReportSo = []entity.ReportSoRes{
	{TotalLoss: 50000, ProductMin: "Oreo", TotalMin: 10000, ProductMax: "SilverQueen", TotalMax: 25000},
	{TotalLoss: 50000, ProductMin: "Richeese", TotalMin: 8000, ProductMax: "SilverQueen", TotalMax: 20000},
	{TotalLoss: 50000, ProductMin: "Tango", TotalMin: 8000, ProductMax: "Minyak Goreng", TotalMax: 20000},
}

type reportProductSoRepoMock struct {
	mock.Mock
}

func (r *reportProductSoRepoMock) GetAlDetailSoReport() any {
	args := r.Called()
	if args[0] == nil {
		return "no data"
	}
	return dummyReportSo
}
func (r *reportProductSoRepoMock) GetAllInterimSoReport() any {
	args := r.Called()
	if args == nil {
		return "no data"
	}

	return dummyReportSo
}

type reportProductSoUCTestSuite struct {
	reportProductSoRepoMock *reportProductSoRepoMock
	suite.Suite
}

func (suite *reportProductSoUCTestSuite) TestGetAllReportTrxSo_Success() {
	reportProductSo := NewReportSoUseCase(suite.reportProductSoRepoMock)
	suite.reportProductSoRepoMock.On("GetAlDetailSoReport").Return(dummyReportSo)
	value := reportProductSo.FindAllReportSoDetail()
	assert.Equal(suite.T(), value, dummyReportSo)

}
func (suite *reportProductSoUCTestSuite) TestGetAllReportTrxSo_Failed() {
	reportProductSo := NewReportSoUseCase(suite.reportProductSoRepoMock)
	suite.reportProductSoRepoMock.On("GetAlDetailSoReport").Return(nil)
	value := reportProductSo.FindAllReportSoDetail()
	assert.Equal(suite.T(), value, "no data")
}
func (suite *reportProductSoUCTestSuite) TestGetAllInterimReportTrxSo_Success() {
	reportProductSo := NewReportSoUseCase(suite.reportProductSoRepoMock)
	suite.reportProductSoRepoMock.On("GetAllInterimSoReport").Return(dummyReportSo)
	value := reportProductSo.FindAllInterimSoReport()
	assert.Equal(suite.T(), value, dummyReportSo)

}
func (suite *reportProductSoUCTestSuite) TestGetAllInterimReportTrxSo_Failed() {
	reportProductSo := NewReportSoUseCase(suite.reportProductSoRepoMock)
	suite.reportProductSoRepoMock.On("GetAllInterimSoReport").Return(nil)
	value := reportProductSo.FindAllInterimSoReport()
	assert.Equal(suite.T(), value, "no data")
}

// ==================================Setup ===================================================
func (suite *reportProductSoUCTestSuite) SetupTest() {
	suite.reportProductSoRepoMock = new(reportProductSoRepoMock)
}

func TestReportProductSoUCTestSuite(t *testing.T) {
	suite.Run(t, new(reportProductSoUCTestSuite))
}
