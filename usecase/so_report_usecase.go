package usecase

import "go_inven_ctrl/repository"

type ReporSoUsecase interface {
	FindAllInterimSoReport() any
	FindAllReportSoDetail() any
}

type reportSoUsecase struct {
	reportSoRepo repository.ReportSoRepo
}

func (u *reportSoUsecase) FindAllInterimSoReport() any {
	return u.reportSoRepo.GetAllInterimSoReport()
}

func (u *reportSoUsecase) FindAllReportSoDetail() any {
	return u.reportSoRepo.GetAlDetailSoReport()
}

func NewReportSoUseCase(reportSo repository.ReportSoRepo) ReporSoUsecase {
	return &reportSoUsecase{
		reportSoRepo: reportSo,
	}
}
