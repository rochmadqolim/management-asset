package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type TrxInSoUsecase interface {
	EnrollInsertTrxSo(enrollTrxInSt *entity.TrxInSo) string
	ReportConfirmation() string
	ReportInterim() string
}

type trxInSoUsecase struct {
	trxInSoRepo repository.TrxInSoRepo
}

func (tx *trxInSoUsecase) EnrollInsertTrxSo(enrollTrxInSo *entity.TrxInSo) string {
	return tx.trxInSoRepo.EnrollInsertTrxInSo(enrollTrxInSo)
}

func (tx *trxInSoUsecase) ReportConfirmation() string {
	return tx.trxInSoRepo.EnrollInsertReportConfirm()
}
func (tx *trxInSoUsecase) ReportInterim() string {
	return tx.trxInSoRepo.EnrollInsertReportInterim()
}

func NewTrxInSoUsecase(trxInSo repository.TrxInSoRepo) TrxInSoUsecase {
	return &trxInSoUsecase{
		trxInSoRepo: trxInSo,
	}
}
