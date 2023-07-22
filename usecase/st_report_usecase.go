package usecase

import (
	"go_inven_ctrl/repository"
)

type ReportTrxStUsecase interface {
	FindAllReportTrxSt() any
	FindByReportTrxProductStId(id string) any
	FindByDateReportTrxSt(date string) any
}

type reportTrxStUsecase struct {
	reportTrxStRepo repository.ReportTrxStRepo
}

func (u *reportTrxStUsecase) FindAllReportTrxSt() any {
	return u.reportTrxStRepo.GetAllReportTrxSt()
}

func (u *reportTrxStUsecase) FindByReportTrxProductStId(id string) any {
	return u.reportTrxStRepo.GetByReportTrxProductStId(id)
}

func (u *reportTrxStUsecase) FindByDateReportTrxSt(date string) any {
	return u.reportTrxStRepo.GetByDateReportTrxSt(date)
}

func NewReportTrxStUsecase(reportTrx repository.ReportTrxStRepo) ReportTrxStUsecase {
	return &reportTrxStUsecase{
		reportTrxStRepo: reportTrx,
	}
}
