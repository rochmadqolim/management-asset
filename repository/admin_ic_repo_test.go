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

var dummyAdminIc = []entity.AdminIc{
	{
		ID:       "1",
		Name:     "dimas",
		Email:    "dimas@mail.com",
		Password: "12345",
		Phone:    "08123456789",
		Photo:    "Screenshot (598).png",
	},
	{
		ID:       "2",
		Name:     "Siti",
		Email:    "siti@mail.com",
		Password: "22345",
		Phone:    "08223456789",
		Photo:    "Screenshot (599).png",
	},
	{
		ID:       "3",
		Name:     "Aliyya",
		Email:    "aliyya@mail.com",
		Password: "32345",
		Phone:    "08323456789",
		Photo:    "Screenshot (600).png",
	},
}

var dummyAdminIcRes = []entity.AdminIc{
	{
		ID:    "1",
		Name:  "dimas",
		Email: "dimas@mail.com",
		Phone: "08123456789",
		Photo: "photo.jpg",
	},
	{
		ID:    "2",
		Name:  "Siti",
		Email: "siti@mail.com",
		Phone: "08223456789",
		Photo: "photo.jpg",
	},
	{
		ID:    "3",
		Name:  "Aliyya",
		Email: "aliyya@mail.com",
		Phone: "08323456789",
		Photo: "photo.jpg",
	},
}

type AdminIcRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *AdminIcRepoTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})

	for _, v := range dummyAdminIcRes {
		rows.AddRow(v.ID, v.Name, v.Email, v.Phone, v.Photo)
	}

	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM ic_team").WillReturnRows(rows)

	repo := NewAdminIcRepo(suite.mockDb)
	expected := dummyAdminIcRes
	actual := repo.GetAll().([]entity.AdminIc)

	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), len(dummyAdminIc), len(actual))
	assert.Equal(suite.T(), "2", actual[1].ID)
}

func (suite *AdminIcRepoTestSuite) TestGetAllScan_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})
	rows.AddRow(nil, "dimas", "dimas@mail.com", "08123456789", "photo.jpg")

	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM ic_team").WillReturnRows(rows)

	repo := NewAdminIcRepo(suite.mockDb)
	actual := repo.GetAll()

	expected := []entity.AdminIc{
		{ID: "", Name: "", Email: "", Phone: "", Photo: ""},
	}
	assert.Equal(suite.T(), expected, actual)
}

func (suite *AdminIcRepoTestSuite) TestGetAll_Empty() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})
	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM ic_team").WillReturnRows(rows)

	repo := NewAdminIcRepo(suite.mockDb)
	actual := repo.GetAll()

	assert.Equal(suite.T(), "no data", actual)
}

func (suite *AdminIcRepoTestSuite) TestGetById_Success() {
	expectedPhoto := "Screenshot (600).png"

	suite.mockSql.ExpectQuery("SELECT photo FROM ic_team WHERE id = ?").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"photo"}).AddRow(expectedPhoto))

	repo := NewAdminIcRepo(suite.mockDb)
	actualPhoto := repo.GetById("1")

	assert.Equal(suite.T(), expectedPhoto, actualPhoto)
}

func (suite *AdminIcRepoTestSuite) TestGetById_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM ic_team WHERE id = ?").WithArgs("admin ic not found").WillReturnError(errors.New("Failed to get photo"))

	repo := NewAdminIcRepo(suite.mockDb)
	actualPhoto := repo.GetById("invalid-id")

	assert.NotNil(suite.T(), actualPhoto)
	assert.Equal(suite.T(), "admin ic not found", actualPhoto)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcCreate_Success() {
	newAdminIc := dummyAdminIc[0]
	suite.mockSql.ExpectExec("INSERT INTO ic_team\\(id, name, email, phone, photo, password\\) VALUES").WithArgs(
		newAdminIc.ID,
		newAdminIc.Name,
		newAdminIc.Email,
		newAdminIc.Phone,
		newAdminIc.Photo,
		newAdminIc.Password,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	adminIcRepository := NewAdminIcRepo(suite.mockDb)
	actual := adminIcRepository.Create(&newAdminIc)
	assert.Equal(suite.T(), "user craeted successfully", actual)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcCreate_Failed() {
	newAdminIc := dummyAdminIc[0]
	suite.mockSql.ExpectExec("INSERT INTO ic_team\\(id, name, email, phone, photo, password\\) VALUES").WillReturnError(errors.New("failed to create user"))

	adminIcRepository := NewAdminIcRepo(suite.mockDb)
	actual := adminIcRepository.Create(&newAdminIc)
	assert.Equal(suite.T(), "failed to create user", actual)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcUpdate_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo", "password"}).AddRow(dummyAdminIc[0].ID, dummyAdminIc[0].Name, dummyAdminIc[0].Email, dummyAdminIc[0].Photo, dummyAdminIc[0].Phone, dummyAdminIc[0].Password)

	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo, password FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE ic_team SET id = \\$1, name = \\$2, email = \\$3, phone = \\$4, photo = \\$5, password = \\$6 WHERE id = \\$7").
		WithArgs(dummyAdminIc[0].ID, dummyAdminIc[0].Name, dummyAdminIc[0].Email, "67890", dummyAdminIc[0].Phone, dummyAdminIc[0].Photo, dummyAdminIc[0].ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewAdminIcRepo(suite.mockDb)
	updatedAdminIc := &entity.AdminIc{
		ID:       dummyAdminIc[0].ID,
		Name:     dummyAdminIc[0].Name,
		Email:    dummyAdminIc[0].Email,
		Password: "67890",
		Phone:    dummyAdminIc[0].Phone,
		Photo:    dummyAdminIc[0].Photo,
	}
	expected := fmt.Sprintf("admin ic with id %v updated successfully", updatedAdminIc.ID)
	actual := repo.Update(updatedAdminIc)
	if actual != expected {
		err := repo.GetById(dummyAdminIc[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcUpdate_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo, password FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnError(errors.New("admin not found"))

	repo := NewAdminIcRepo(suite.mockDb)
	expected := "admin not found"
	actual := repo.Update(&dummyAdminIc[0])

	assert.Equal(suite.T(), expected, actual)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcDelete_Success() {
	suite.mockSql.ExpectQuery("SELECT photo FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(dummyAdminIc[0].ID))

	suite.mockSql.ExpectExec("DELETE FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewAdminIcRepo(suite.mockDb)
	expected := fmt.Sprintf("admin ic with id %s deleted successfully", dummyAdminIc[0].ID)
	actual := repo.Delete(dummyAdminIc[0].ID)
	if actual != expected {
		err := repo.GetById(dummyAdminIc[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *AdminIcRepoTestSuite) TestAdminIcDelete_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewAdminIcRepo(suite.mockDb)
	expected := "failed to delete admin ic"
	actual := repo.Delete(dummyAdminIc[0].ID)

	assert.Equal(suite.T(), expected, actual)

}

func (suite *AdminIcRepoTestSuite) TestAdminIcDelete_NotFound() {
	suite.mockSql.ExpectQuery("SELECT photo FROM ic_team WHERE id = ?").WithArgs(dummyAdminIc[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewAdminIcRepo(suite.mockDb)
	expected := "failed to delete admin ic"
	actual := repo.Delete(dummyAdminIc[0].ID)

	assert.Equal(suite.T(), expected, actual)
}

func (suite *AdminIcRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *AdminIcRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestAdminIcRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AdminIcRepoTestSuite))
}
