package repository

import (
	"database/sql"
	"go_inven_ctrl/entity"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TrxSoRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

var dummyTrxSo = []entity.TrxInSo{
	{
		ID: 1, ProductStSoId: "pst1", Stock: 20,
	},
	{
		ID: 1, ProductStSoId: "pst1", Stock: -10,
	},
	{
		ID: 1, ProductStSoId: "pst2", Stock: 20,
	},
	{
		ID: 1, ProductStSoId: "pst4", Stock: 35,
	},
}

func (suite *TrxSoRepoTestSuite) TestEnrollTrxSo_Failed() {
	product := dummyTrxSo[0]

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("INSERT INTO trx_so \\(product_so_st_id,stock\\) values").WithArgs(product.ProductStSoId, product.Stock).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mockSql.ExpectQuery("select sum \\(pso.stock \\+ txs.stock\\) from trx_so as txs join product_so as pso on txs.product_so_st_id = pso.product_st_id where txs.product_so_st_id = \\$1;").
		WithArgs(product.ProductStSoId)

	suite.mockSql.ExpectExec("update product_so set stock = \\$1 where product_st_id = \\$2").WithArgs(product.Stock, product.ProductStSoId)
	suite.mockSql.ExpectExec(`update product_so set diff_stock = product_so.stock - product_st.stock, diff_price = (product_st.price * product_so.stock)-(product_st.price * product_st.stock) from product_st where product_so.product_st_id = product_st.id;`)
	suite.mockSql.ExpectExec("delete from trx_so where product_so_st_id = \\$1").WithArgs(product.ID)
	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertTrxInSo(&product)
	assert.Equal(suite.T(), str, "product not found")

}

func (suite *TrxSoRepoTestSuite) TestEnrollTrxSo_Failed2() {
	product := dummyTrxSo[0]

	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertTrxInSo(&product)
	assert.Equal(suite.T(), str, "Transaction rollback")

}

// ====================================================Interim confirm=================================================================
func (suite *TrxSoRepoTestSuite) TestEnrollTrxSoInterim_Failed() {
	product := dummyTrxSo[0]
	var totalLoss int

	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectQuery("select sum\\(diff_price\\) from product_so;").WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(
		totalLoss,
	))

	suite.mockSql.ExpectExec("update product_so set stock = \\$1 where product_st_id = \\$2").WithArgs(product.Stock, product.ProductStSoId)
	suite.mockSql.ExpectExec(`update product_so set diff_stock = product_so.stock - product_st.stock, diff_price = (product_st.price * product_so.stock)-(product_st.price * product_st.stock) from product_st where product_so.product_st_id = product_st.id;`)
	suite.mockSql.ExpectExec("delete from trx_so where product_so_st_id = \\$1").WithArgs(product.ID)
	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertReportInterim()
	assert.Equal(suite.T(), str, "product not found")

}
func (suite *TrxSoRepoTestSuite) TestEnrollTrxSoInterim_Failed2() {
	product := dummyTrxSo[0]
	var totalLoss int

	suite.mockSql.ExpectQuery("select sum\\(diff_price\\) from product_so;").WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(
		totalLoss,
	))

	suite.mockSql.ExpectExec("update product_so set stock = \\$1 where product_st_id = \\$2").WithArgs(product.Stock, product.ProductStSoId)
	suite.mockSql.ExpectExec(`update product_so set diff_stock = product_so.stock - product_st.stock, diff_price = (product_st.price * product_so.stock)-(product_st.price * product_st.stock) from product_st where product_so.product_st_id = product_st.id;`)
	suite.mockSql.ExpectExec("delete from trx_so where product_so_st_id = \\$1").WithArgs(product.ID)
	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertReportInterim()
	assert.Equal(suite.T(), str, "Transaction rollback")

}

// =========================================Confirm ========================================================
func (suite *TrxSoRepoTestSuite) TestEnrollTrxSoConfirm_Failed2() {

	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertReportInterim()
	assert.Equal(suite.T(), str, "Transaction rollback")

}
func (suite *TrxSoRepoTestSuite) TestEnrollTrxSoConfirm_Failed() {

	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectCommit()

	soRepo := NewTrxInSoRepo(suite.mockDb)
	str := soRepo.EnrollInsertReportConfirm()
	assert.Equal(suite.T(), str, "product not found")

}

//================================================================confirm====================================================

// ============== Before each =====================
func (suite *TrxSoRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("An Error When opening a stub database Connection", err)
	}

	suite.mockDb = mockDb
	suite.mockSql = mockSql
}
func TestTrxSoTestSuite(t *testing.T) {
	suite.Run(t, new(TrxSoRepoTestSuite))
}
func (suite *TrxSoRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}
