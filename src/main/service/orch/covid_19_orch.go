package orch

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/covid19-service/src/main/client"
	"github.com/covid19-service/src/main/dtos"
	"github.com/covid19-service/src/main/service"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Covid19Orch struct {
	RedisClient    *redis.Client
	GoogleClient   client.GoogleClient
	Covid19Service service.Covid19Service
}

func NewCovid19Orch(redisClient *redis.Client, googleClient client.GoogleClient, covid19Service service.Covid19Service) *Covid19Orch {
	return &Covid19Orch{
		RedisClient:    redisClient,
		GoogleClient:   googleClient,
		Covid19Service: covid19Service,
	}
}
func (orch *Covid19Orch) UpdateCovid19Case() (interface{}, error) {
	covid19Cases, err := orch.GoogleClient.GetIndiaCovidCases()
	if err == nil {
		err = orch.Covid19Service.UpdateCovid19Case(covid19Cases)
		if err == nil {
			return covid19Cases, nil
		}
	}
	return nil, err
}

func (orch *Covid19Orch) GetCovid19Data(request dtos.GetCovid19CaseByPlaceRequest) (interface{}, error) {
	var res []dtos.GetCovid19CaseByPlaceResponse
	placeDetails, err := orch.GoogleClient.GetStateNameFromLatLng(request.Lat, request.Lng, request.ApiKey)
	if err == nil {

		state := placeDetails.Region
		state = strings.Replace(state, "&", "and", -1)

		val, err := orch.RedisClient.Get(state).Result()
		if err == nil {
			logrus.Info("[GetCovid19Data] Getting data from redis ! val : ", val)
			_ = json.Unmarshal([]byte(val), &res)
			return res, nil
		} else {
			logrus.Error("[GetCovid19Data] Error in getting redis data : ", err)
		}

		covidCases, err := orch.Covid19Service.GetCovidCaseByPlace(state)
		if err == nil {
			valueByte, _ := json.Marshal(covidCases)
			err = orch.RedisClient.Set(state, string(valueByte), time.Minute*30).Err()
			if err != nil {
				logrus.Error("[GetCovid19Data] Error in adding data in redis : ", err)
			}

			return covidCases, nil
		}

		return nil, err

	}

	return nil, err
}
