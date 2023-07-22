package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummySeller = []entity.Storeteam{
	{
		ID:       "1",
		Name:     "Seller 1",
		Email:    "seller1@mail.com",
		Password: "12345",
		Phone:    "087432473247",
		Photo:    "photo.jpg",
	},
	{
		ID:       "2",
		Name:     "Seller 2",
		Email:    "seller2@mail.com",
		Password: "23456",
		Phone:    "0822222",
		Photo:    "photo.jpg",
	},
	{
		ID:       "3",
		Name:     "Seller 3",
		Email:    "seller3@mail.com",
		Password: "34567",
		Phone:    "0877789988",
		Photo:    "photo.jpg",
	},
}

type repoMockStTeam struct {
	mock.Mock
}

type StTeamUsecaseTestSuite struct {
	repoMockStTeam *repoMockStTeam
	suite.Suite
}

func (r *repoMockStTeam) GetAll() any {
	arg := r.Called()
	if arg.Get(0) == nil {
		return []entity.Storeteam{}
		// return "no data"
	}
	return arg.Get(0).([]entity.Storeteam)
}

func (r *repoMockStTeam) GetById(id string) any {
	arg := r.Called(id)
	if arg.Get(0) == nil {
		return "seller not found"
	}
	return arg.Get(0)
}

func (r *repoMockStTeam) Create(newSeller *entity.Storeteam) string {
	arg := r.Called(newSeller)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "seller created"
}

func (r *repoMockStTeam) Update(seller *entity.Storeteam) string {
	arg := r.Called(seller)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "seller updated"
}

func (r *repoMockStTeam) Delete(id string) string {
	arg := r.Called(id)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "seller deleted"
}

func (suite *StTeamUsecaseTestSuite) TestRegister_Success() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Create", &dummySeller[0]).Return("seller created")

	sellerSt := stTeamUc.Register(&dummySeller[0])

	assert.Equal(suite.T(), "seller created", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestRegister_Failed() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Create", &dummySeller[0]).Return("failed to create seller")

	sellerSt := stTeamUc.Register(&dummySeller[0])

	assert.Equal(suite.T(), "failed to create seller", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestFindSellers_Success() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("GetAll").Return(dummySeller)

	sellerSt := stTeamUc.FindSellers()
	sellerSts := sellerSt.([]entity.Storeteam)

	assert.Equal(suite.T(), dummySeller, sellerSt)
	assert.Equal(suite.T(), len(dummySeller), len(sellerSts))
}
func (suite *StTeamUsecaseTestSuite) TestFindSellers_Failed() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("GetAll").Return([]entity.Storeteam{})
	// suite.repoMockStTeam.On("GetAll").Return("no data")

	sellerSt := stTeamUc.FindSellers()
	sellerSts := sellerSt.([]entity.Storeteam)

	// assert.Equal(suite.T(),"no data", sellerSt)
	assert.Equal(suite.T(), 0, len(sellerSts))
	assert.Empty(suite.T(), sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestFindSellerById_Success() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("GetById", "1").Return(dummySeller[0].Photo)

	sellerSt := stTeamUc.FindSellerById("1")

	assert.Equal(suite.T(), dummySeller[0].Photo, sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestFindSellerById_Failed() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("GetById", "5").Return("no data")

	sellerSt := stTeamUc.FindSellerById("5")

	assert.Equal(suite.T(), "no data", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestEdit_Success() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Update", &dummySeller[0]).Return("admin updated")

	sellerSt := stTeamUc.Edit(&dummySeller[0])

	assert.Equal(suite.T(), "admin updated", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestEdit_Failed() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Update", &dummySeller[0]).Return("failed to update seller")

	sellerSt := stTeamUc.Edit(&dummySeller[0])

	assert.Equal(suite.T(), "failed to update seller", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestUnreg_Success() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Delete", "1").Return("admin deleted")

	sellerSt := stTeamUc.Unreg("1")

	assert.Equal(suite.T(), "admin deleted", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) TestUnreg_Failed() {
	stTeamUc := NewStoreteamUsecase(suite.repoMockStTeam)
	suite.repoMockStTeam.On("Delete", "5").Return("failed to delete seller")

	sellerSt := stTeamUc.Unreg("5")

	assert.Equal(suite.T(), "failed to delete seller", sellerSt)
}

func (suite *StTeamUsecaseTestSuite) SetupTest() {
	suite.repoMockStTeam = new(repoMockStTeam)
}

func TestStTeamUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(StTeamUsecaseTestSuite))
}
