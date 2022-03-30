package impl

import (
	"strings"

	"github.com/covid19-service/src/main/dal/models"
	"github.com/covid19-service/src/main/dal/repos"
	"github.com/covid19-service/src/main/dtos"
	"github.com/covid19-service/src/main/service"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type Covid19Impl struct {
	Covid19Repo *repos.Covid19Repo
}

func NewCovid19Impl(Covid19Repo *repos.Covid19Repo) service.Covid19Service {
	return &Covid19Impl{Covid19Repo: Covid19Repo}
}

func (service *Covid19Impl) UpdateCovid19Case(covid19Case *dtos.CovidCase) error {
	covid19DataState := covid19Case.Data.Regional
	var (
		err               error
		covid19CaseModels []*models.Covid19
		covid19Cases      []models.Covid19
	)

	for _, state := range covid19DataState {
		covid19Cases, err = service.Covid19Repo.GetByStates([]string{state.StateName})
		if err != nil || len(covid19Cases) == 0 {
			id := strings.Replace(uuid.New().String(), "-", "", -1)
			covid19CaseModel := models.NewCovid19(id, state.StateName, covid19Case.LastOriginUpdate, state.Deaths, state.Recovered, state.TotalConfirmed)
			covid19CaseModels = append(covid19CaseModels, covid19CaseModel)
			continue
		}
		covid19Cases[0].UpdatedAt = covid19Case.LastOriginUpdate
		covid19Cases[0].Deaths = state.Deaths
		covid19Cases[0].Recovered = state.Recovered
		covid19Cases[0].ConfirmedCovidCase = state.TotalConfirmed
		err = service.Covid19Repo.UpsertByPlace(&covid19Cases[0])
		if err != nil {
			logrus.Error("Error in updating data for state : {} Err : {} ", state.StateName, err)
		}
	}
	covid19CaseIndia := covid19Case.Data.Summary
	covid19Cases, err = service.Covid19Repo.GetByStates([]string{"India"})
	if err == nil && len(covid19Cases) > 0 {
		covid19Cases[0].UpdatedAt = covid19Case.LastOriginUpdate
		covid19Cases[0].Deaths = covid19CaseIndia.Deaths
		covid19Cases[0].Recovered = covid19CaseIndia.Recovered
		covid19Cases[0].ConfirmedCovidCase = covid19CaseIndia.Total
		err = service.Covid19Repo.UpsertByPlace(&covid19Cases[0])
		if err != nil {
			logrus.Error("Error in updating data for India  Err : ", err)
		}
	} else {
		id := strings.Replace(uuid.New().String(), "-", "", -1)
		covid19CaseModel := models.NewCovid19(id, "India", covid19Case.LastOriginUpdate, covid19CaseIndia.Deaths, covid19CaseIndia.Recovered, covid19CaseIndia.Total)
		covid19CaseModels = append(covid19CaseModels, covid19CaseModel)
	}
	if len(covid19CaseModels) > 0 {
		_, err = service.Covid19Repo.BulkInsert(covid19CaseModels)
		if err != nil {
			logrus.Error("[Covid19Impl] UpdateCovid19Case  error : ", err)
		}
	}
	return nil
}

func (service *Covid19Impl) GetCovidCaseByPlace(state string) ([]dtos.GetCovid19CaseByPlaceResponse, error) {
	var response []dtos.GetCovid19CaseByPlaceResponse
	states := []string{"India"}
	states = append(states, state)
	covid19Cases, err := service.Covid19Repo.GetByStates(states)
	if err == nil {
		for _, covid19Case := range covid19Cases {
			response = append(response, dtos.GetCovid19CaseByPlaceResponse{
				Place:              covid19Case.PlaceName,
				Deaths:             covid19Case.Deaths,
				Recovered:          covid19Case.Recovered,
				ConfirmedCovidCase: covid19Case.ConfirmedCovidCase,
				UpdatedAt:          covid19Case.UpdatedAt,
			})
		}
		return response, nil
	}
	logrus.Error("[Covid19Impl] GetCovidCaseByPlace  error : ", err)
	return nil, err
}
