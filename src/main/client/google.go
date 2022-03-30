package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/covid19-service/src/main/dtos"
)

type GoogleClient struct {
	client *http.Client
}

func NewGoogleClient() GoogleClient {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: time.Duration(100) * time.Second,
	}
	return GoogleClient{
		client: client,
	}
}

func (gc *GoogleClient) GetIndiaCovidCases() (*dtos.CovidCase, error) {
	url := "https://api.rootnet.in/covid19-in/stats/latest"
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		var res *http.Response
		res, err = gc.client.Do(req)
		if err == nil {
			defer res.Body.Close()
			if res.StatusCode == 200 {
				var body []byte
				body, err = ioutil.ReadAll(res.Body)
				if err == nil {
					var data dtos.CovidCase
					err = json.Unmarshal(body, &data)
					if err == nil {
						return &data, nil
					}
				}
			}
		}
	}
	logrus.Error("[GoogleClient] GetIndiaCovidCases Error : ", err)
	return nil, err
}

//key : 0045b01eda7d5aaf7f84ee668863e35f
func (gc *GoogleClient) GetStateNameFromLatLng(lat string, lng string, apiKey string) (dtos.ReverseGeoCodingData, error) {
	url := "http://api.positionstack.com/v1/reverse?access_key=" + apiKey + "&query=" + lat + "," + lng
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		var res *http.Response
		res, err = gc.client.Do(req)
		if err == nil {
			defer res.Body.Close()
			if res.StatusCode == 200 {
				var body []byte
				body, err = ioutil.ReadAll(res.Body)
				if err == nil {
					var data dtos.ReverseGeoCodingResponse
					err = json.Unmarshal(body, &data)
					if err == nil {
						return data.ReverseGeoCodingData[0], nil
					}
				}
			}
		}
	}
	logrus.Error("[GoogleClient] GetStateNameFromLatLng  Error : ", err, "for lat : ", lat, "lng : ", lng)
	return dtos.ReverseGeoCodingData{}, err
}
