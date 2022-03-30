package service

import "github.com/covid19-service/src/main/dtos"

type Covid19Service interface {
	UpdateCovid19Case(covid19Case *dtos.CovidCase) error
	GetCovidCaseByPlace(state string) ([]dtos.GetCovid19CaseByPlaceResponse, error)
}
