package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyReportSt = []entity.ReportTrxSt{
	{ProductStId: "pst1", StockIn: 100, ProductName: "SilverKing", Act: "input", LastStock: 0, CreatedAt: "2023-04-01"},
	{ProductStId: "pst1", StockIn: 20, ProductName: "SilverKing", Act: "input", LastStock: 150, CreatedAt: "2023-04-01"},
	{ProductStId: "pst1", StockIn: 50, ProductName: "SilverKing", Act: "input", LastStock: 100, CreatedAt: "2023-04-02"},
}

type reportProductStRepoMock struct {
	mock.Mock
}

func (r *reportProductStRepoMock) GetAllReportTrxSt() any {
	args := r.Called()
	if args[0] == nil {
		return "no data"
	}

	return dummyReportSt
}
func (r *reportProductStRepoMock) GetByReportTrxProductStId(id string) any {
	args := r.Called(id)
	if args[0] == nil {
		return "no data or product id not found"
	}

	return dummyReportSt
}
func (r *reportProductStRepoMock) GetByDateReportTrxSt(date string) any {
	args := r.Called(date)
	if args[0] == nil {
		return "no data or date product not found"
	}
	return dummyReportSt
}

type reportProductStUCTestSuite struct {
	reportProductStRepoMock *reportProductStRepoMock
	suite.Suite
}

func (suite *reportProductStUCTestSuite) TestGetAllReportTrxSt_Success() {
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetAllReportTrxSt").Return(dummyReportSt)
	value := reportProductSt.FindAllReportTrxSt()
	assert.Equal(suite.T(), value, dummyReportSt)
}
func (suite *reportProductStUCTestSuite) TestGetAllReportTrxSt_Failed() {
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetAllReportTrxSt").Return(nil)
	value := reportProductSt.FindAllReportTrxSt()
	assert.Equal(suite.T(), value, "no data")
}

// ==========================================by date============================================
func (suite *reportProductStUCTestSuite) TestGetAlByDatelReportTrxSt_Success() {
	date := "2023-04-01"
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetByDateReportTrxSt", date).Return(dummyReportSt)
	value := reportProductSt.FindByDateReportTrxSt(date)
	assert.Equal(suite.T(), value, dummyReportSt)
}
func (suite *reportProductStUCTestSuite) TestGetAllByDateReportTrxSt_Failed() {
	date := "2023-04-01"
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetByDateReportTrxSt", date).Return(nil)
	value := reportProductSt.FindByDateReportTrxSt(date)
	assert.Equal(suite.T(), value, "no data or date product not found")
}

// ==========================================by StId============================================
func (suite *reportProductStUCTestSuite) TestGetAlByStIdlReportTrxSt_Success() {
	id := "pst1"
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetByReportTrxProductStId", id).Return(dummyReportSt)
	value := reportProductSt.FindByReportTrxProductStId(id)
	assert.Equal(suite.T(), value, dummyReportSt)
}
func (suite *reportProductStUCTestSuite) TestGetAllByStIdeReportTrxSt_Failed() {
	id := "pst1"
	reportProductSt := NewReportTrxStUsecase(suite.reportProductStRepoMock)
	suite.reportProductStRepoMock.On("GetByReportTrxProductStId", id).Return(nil)
	value := reportProductSt.FindByReportTrxProductStId(id)
	assert.Equal(suite.T(), value, "no data or product id not found")
}

// ==================================Setup ===================================================
func (suite *reportProductStUCTestSuite) SetupTest() {
	suite.reportProductStRepoMock = new(reportProductStRepoMock)
}

func TestReportProductStUCTestSuite(t *testing.T) {
	suite.Run(t, new(reportProductStUCTestSuite))
}
