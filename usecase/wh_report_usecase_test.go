package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"go_inven_ctrl/entity"
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
	{
		ProductWhId: "P002",
		Stock:       200,
		ProductName: "Yakult",
		Act:         "output",
		LastStock:   2000,
		CreatedAt:   "2023-01-01",
	},
}

type repoMockWhReport struct {
	mock.Mock
}

type WhReportRepoTestSuite struct {
	repoMockWhReport *repoMockWhReport
	suite.Suite
}

func (r *repoMockWhReport) GetAllReportTrxWh() any {
	args := r.Called()
	if args.Get(0) == nil {
		return []entity.ReportWh{}
	}
	return args.Get(0).([]entity.ReportWh)
}

func (r *repoMockWhReport) GetByIdReportTrxWh(id string) any {
	args := r.Called(id)
	if args.Get(0) == nil {
		return "transaction not found"
	}
	return args.Get(0)
}

func (r *repoMockWhReport) GetByDateReportTrxWh(date string) any {
	args := r.Called(date)
	if args.Get(0) == nil {
		return "transaction not found"
	}
	return args.Get(0)
}

func (suite *WhReportRepoTestSuite) TestFindAllReportTrxWh_Success() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetAllReportTrxWh").Return(dummyReportWh)

	whReport := whReportUc.FindAllReportTrxWh()
	whReports := whReport.([]entity.ReportWh)

	assert.Equal(suite.T(), dummyReportWh, whReport)
	assert.Equal(suite.T(), len(dummyReportWh), len(whReports))
}

func (suite *WhReportRepoTestSuite) TestFindAllReportTrxWh_Failed() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetAllReportTrxWh").Return([]entity.ReportWh{})

	whReport := whReportUc.FindAllReportTrxWh()
	whReports := whReport.([]entity.ReportWh)

	assert.Equal(suite.T(), 0, len(whReports))
	assert.Empty(suite.T(), whReport)
}

func (suite *WhReportRepoTestSuite) TestFindByIdReportTrxWh_Success() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetByIdReportTrxWh", "1").Return(dummyReportWh[0].ProductWhId)

	adminWh := whReportUc.FindByIdReportTrxWh("1")

	assert.Equal(suite.T(), dummyReportWh[0].ProductWhId, adminWh)
}

func (suite *WhReportRepoTestSuite) TestFindByIdReportTrxWh_Failed() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetByIdReportTrxWh", "5").Return("no data")

	adminWh := whReportUc.FindByIdReportTrxWh("5")

	assert.Equal(suite.T(), "no data", adminWh)
}

func (suite *WhReportRepoTestSuite) TestFindByDateReportTrxWh_Success() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetByDateReportTrxWh", "2023-01-01").Return(dummyReportWh[0].CreatedAt)

	adminWh := whReportUc.FindByDateReportTrxWh("2023-01-01")

	assert.Equal(suite.T(), dummyReportWh[0].CreatedAt, adminWh)
}

func (suite *WhReportRepoTestSuite) TestFindByDateReportTrxWh_Failed() {
	whReportUc := NewReportTrxWhUsecase(suite.repoMockWhReport)
	suite.repoMockWhReport.On("GetByDateReportTrxWh", "2000-00-00").Return("no data")

	adminWh := whReportUc.FindByDateReportTrxWh("2000-00-00")

	assert.Equal(suite.T(), "no data", adminWh)
}

func (suite *WhReportRepoTestSuite) SetupTest() {
	suite.repoMockWhReport = new(repoMockWhReport)
}

func TestTrxWhRepoTestSuite(t *testing.T) {
	suite.Run(t, new(WhReportRepoTestSuite))
}
