package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go_inven_ctrl/entity"

	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummySellernil = []entity.Storeteam{}
var dummySeller = []entity.Storeteam{
	{
		ID:    "1",
		Name:  "Seller 1",
		Email: "seller1@mail.com",
		Phone: "087432473247",
		Photo: "photo.jpg",
	},
	{
		ID:    "2",
		Name:  "Seller 2",
		Email: "seller2@mail.com",
		Phone: "0822222",
		Photo: "photo.jpg",
	},
	{
		ID:    "3",
		Name:  "Seller 3",
		Email: "seller3@mail.com",
		Phone: "0877789988",
		Photo: "photo.jpg",
	},
}

type StoreTeamRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *StoreTeamRepoTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})
	query := "SELECT id, name, email, phone, photo FROM st_team"
	for _, v := range dummySeller {
		rows.AddRow(v.ID, v.Name, v.Email, v.Phone, v.Photo)
	}

	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)

	repo := NewStoreteamRepo(suite.mockDb)
	expected := dummySeller
	actual := repo.GetAll().([]entity.Storeteam)

	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), len(dummySeller), len(actual))
	assert.Equal(suite.T(), "2", actual[1].ID)
}

func (suite *StoreTeamRepoTestSuite) TestGetAll_Failed() {

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})

	query := "SELECT id, name, email, phone, photo FROM st_team"
	for _, v := range dummySellernil {
		rows.AddRow(v.ID, v.Name, v.Email, v.Phone, v.Photo)
	}

	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)

	repo := NewStoreteamRepo(suite.mockDb)

	expected := "no data"
	actual := repo.GetAll()

	assert.Equal(suite.T(), expected, actual)

}

func (suite *StoreTeamRepoTestSuite) TestGetById_Success() {
	expectedPhoto := "Seller 1.png"

	suite.mockSql.ExpectQuery("SELECT photo FROM st_team WHERE id = ?").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"photo"}).AddRow(expectedPhoto))

	repo := NewStoreteamRepo(suite.mockDb)
	actualPhoto := repo.GetById("1")

	assert.Equal(suite.T(), expectedPhoto, actualPhoto)
}

func (suite *StoreTeamRepoTestSuite) TestGetById_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM st_team WHERE id = ?").WithArgs("seller not found").WillReturnError(errors.New("Failed to get photo"))

	repo := NewStoreteamRepo(suite.mockDb)
	actualPhoto := repo.GetById("invalid-id")

	assert.NotNil(suite.T(), actualPhoto)
	assert.Equal(suite.T(), "seller not found", actualPhoto)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamCreate_Success() {
	newStoreteam := dummySeller[0]
	suite.mockSql.ExpectExec("INSERT INTO st_team\\(id, name, email, password, phone, photo\\) VALUES").WithArgs(
		newStoreteam.ID,
		newStoreteam.Name,
		newStoreteam.Email,
		newStoreteam.Password,
		newStoreteam.Phone,
		newStoreteam.Photo,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	StoreteamRepository := NewStoreteamRepo(suite.mockDb)
	actual := StoreteamRepository.Create(&newStoreteam)
	assert.Equal(suite.T(), "seller craeted successfully", actual)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamCreate_Failed() {
	newStoreteam := dummySeller[0]
	suite.mockSql.ExpectExec("INSERT INTO st_team\\(id, name, email, password, phone, photo\\) VALUES").WillReturnError(errors.New("failed to create seller"))

	StoreteamRepository := NewStoreteamRepo(suite.mockDb)
	actual := StoreteamRepository.Create(&newStoreteam)
	assert.Equal(suite.T(), "failed to create seller", actual)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamUpdate_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "phone", "photo"}).AddRow(dummySeller[0].ID, dummySeller[0].Name, dummySeller[0].Email, dummySeller[0].Password, dummySeller[0].Photo, dummySeller[0].Phone)

	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE st_team SET id = \\$1, name = \\$2, email = \\$3,  password = \\$4, phone = \\$5, photo = \\$6 WHERE id = \\$7").
		WithArgs(dummySeller[0].ID, dummySeller[0].Name, dummySeller[0].Email, "12345", dummySeller[0].Phone, dummySeller[0].Photo, dummySeller[0].ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewStoreteamRepo(suite.mockDb)
	updatedStoreteam := &entity.Storeteam{
		ID:       dummySeller[0].ID,
		Name:     dummySeller[0].Name,
		Email:    dummySeller[0].Email,
		Password: "12345",
		Phone:    dummySeller[0].Phone,
		Photo:    dummySeller[0].Photo,
	}
	expected := fmt.Sprintf("seller with id %v updated successfully", updatedStoreteam.ID)
	actual := repo.Update(updatedStoreteam)
	if actual != expected {
		err := repo.GetById(dummySeller[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamUpdate_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnError(errors.New("seller not found"))

	repo := NewStoreteamRepo(suite.mockDb)
	expected := "seller not found"
	actual := repo.Update(&dummySeller[0])

	assert.Equal(suite.T(), expected, actual)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamDelete_Success() {
	suite.mockSql.ExpectQuery("SELECT photo FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(dummySeller[0].ID))

	suite.mockSql.ExpectExec("DELETE FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewStoreteamRepo(suite.mockDb)
	expected := fmt.Sprintf("seller with id %s deleted successfully", dummySeller[0].ID)
	actual := repo.Delete(dummySeller[0].ID)
	if actual != expected {
		err := repo.GetById(dummyAdmin[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *StoreTeamRepoTestSuite) TestStoreteamDelete_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewStoreteamRepo(suite.mockDb)
	expected := "seller not found"
	actual := repo.Delete(dummySeller[0].ID)

	assert.Equal(suite.T(), expected, actual)

}

func (suite *StoreTeamRepoTestSuite) TestStoreteamDelete_NotFound() {
	suite.mockSql.ExpectQuery("SELECT photo FROM st_team WHERE id = ?").WithArgs(dummySeller[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewStoreteamRepo(suite.mockDb)
	expected := "seller not found"
	actual := repo.Delete(dummySeller[0].ID)

	assert.Equal(suite.T(), expected, actual)
}

func (suite *StoreTeamRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *StoreTeamRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestStoreTeamRepoTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTeamRepoTestSuite))
}
