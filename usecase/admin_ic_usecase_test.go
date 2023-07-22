package usecase

import (
	"testing"

	"go_inven_ctrl/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyAdminIc = []entity.AdminIc{
	{
		ID:       "C001",
		Name:     "Dummy Name 1",
		Email:    "Dummy Email 1",
		Phone:    "Dummy Phone 1",
		Photo:    "Dummy Photo 1",
		Password: "Dummy Password 1",
	},
	{
		ID:       "C002",
		Name:     "Dummy Name 2",
		Email:    "Dummy Email 2",
		Phone:    "Dummy Phone 2",
		Photo:    "Dummy Photo 2",
		Password: "Dummy Password 2",
	},
}

var dummyAdminIcRes = []entity.AdminIc{
	{
		ID:    "1",
		Name:  "Dodi",
		Email: "dodi@mail.com",
		Phone: "08123456789",
		Photo: "photo.jpg",
	},
	{
		ID:    "2",
		Name:  "Dika",
		Email: "dika@mail.com",
		Phone: "08223456789",
		Photo: "photo.jpg",
	},
}

type repoMock struct {
	mock.Mock
}

type AdminIcUseCaseTestSuite struct {
	repoMock *repoMock
	suite.Suite
}

func (r *repoMock) GetAll() any {
	args := r.Called()
	if args.Get(0) == nil {
		return []entity.AdminIc{}
	}
	return args.Get(0).([]entity.AdminIc)
}

func (r *repoMock) GetById(id string) any {
	args := r.Called(id)
	if args.Get(0) == nil {
		return "admin ic not found"
	}
	return args.Get(0)
}

func (r *repoMock) Create(newAdminIc *entity.AdminIc) string {
	args := r.Called(newAdminIc)
	if args[0] != nil {
		return args.String(0)
	}
	return "admin ic created"
}

func (r *repoMock) Update(adminIc *entity.AdminIc) string {
	args := r.Called(adminIc)
	if args[0] != nil {
		return args.String(0)
	}
	return "admin ic updated"
}

func (r *repoMock) Delete(id string) string {
	args := r.Called(id)
	if args[0] != nil {
		return args.String(0)
	}
	return "admin ic deleted"
}

func (suite *AdminIcUseCaseTestSuite) TestRegisterAdminIc_Success() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Create", &dummyAdminIc[0]).Return("admin ic created")
	adminIc := adminIcUc.Register(&dummyAdminIc[0])
	assert.Equal(suite.T(), "admin ic created", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestRegisterAdminIc_Failed() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Create", &dummyAdminIc[0]).Return("failed to create admin ic")

	adminIc := adminIcUc.Register(&dummyAdminIc[0])

	assert.Equal(suite.T(), "failed to create admin ic", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestGetAllAdminIc_Success() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("GetAll").Return(dummyAdminIcRes)

	adminIc := adminIcUc.FindAdminIc()
	adminIcs := adminIc.([]entity.AdminIc)

	assert.Equal(suite.T(), dummyAdminIcRes, adminIc)
	assert.Equal(suite.T(), len(dummyAdminIcRes), len(adminIcs))
}

func (suite *AdminIcUseCaseTestSuite) TestGetAllAdminIc_Failed() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("GetAll").Return([]entity.AdminIc{})
	adminIc := adminIcUc.FindAdminIc()
	adminIcs := adminIc.([]entity.AdminIc)
	assert.Equal(suite.T(), 0, len(adminIcs))
	assert.Empty(suite.T(), adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestGetAdminIcById_Success() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("GetById", "1").Return(dummyAdminIcRes[0].Photo)

	adminIc := adminIcUc.FindAdminIcById("1")

	assert.Equal(suite.T(), dummyAdminIcRes[0].Photo, adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestGetAdminIcById_Failed() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("GetById", "5").Return("no data")

	adminIc := adminIcUc.FindAdminIcById("5")

	assert.Equal(suite.T(), "no data", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestEditAdminIc_Success() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Update", &dummyAdminIc[0]).Return("admin updated")

	adminIc := adminIcUc.Edit(&dummyAdminIc[0])

	assert.Equal(suite.T(), "admin updated", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestEditAdminIc_Failed() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Update", &dummyAdminIc[0]).Return("failed to update admin ic")

	adminIc := adminIcUc.Edit(&dummyAdminIc[0])

	assert.Equal(suite.T(), "failed to update admin ic", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestUnregAdminIc_Success() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Delete", "1").Return("admin ic deleted")

	adminIc := adminIcUc.Unreg("1")

	assert.Equal(suite.T(), "admin ic deleted", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) TestUnregAdminIc_Failed() {
	adminIcUc := NewAdminIcUsecase(suite.repoMock)
	suite.repoMock.On("Delete", "5").Return("failed to delete admin ic")

	adminIc := adminIcUc.Unreg("5")

	assert.Equal(suite.T(), "failed to delete admin ic", adminIc)
}

func (suite *AdminIcUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestAdminIcUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AdminIcUseCaseTestSuite))
}
