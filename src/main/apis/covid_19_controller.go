package apis

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/covid19-service/src/main/dtos"
	"github.com/covid19-service/src/main/service/orch"
	"github.com/labstack/echo"
)

type Covid19Controller struct {
	covid19Orch *orch.Covid19Orch
}

func NewCovid19Controller(route *echo.Group, covid19Orch *orch.Covid19Orch) *Covid19Controller {
	controller := Covid19Controller{covid19Orch: covid19Orch}
	v0 := route.Group("/v1/covid")
	{
		v0.POST("/update", controller.UpdateCovidCase)
		v0.GET("/cases", controller.GetCovidCase)
	}
	return &controller
}

func (controller *Covid19Controller) UpdateCovidCase(ctx echo.Context) error {
	logrus.Info("[Covid19Controller] UpdateCovidCase request received *****")

	response, err := controller.covid19Orch.UpdateCovid19Case()
	if err != nil {
		logrus.Error("error : ", err)
		return err
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

func (controller *Covid19Controller) GetCovidCase(ctx echo.Context) error {
	logrus.Info("[Covid19Controller] GetCovidCase request received *****")

	defer ctx.Request().Body.Close()

	var request dtos.GetCovid19CaseByPlaceRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		logrus.Fatal("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	response, err := controller.covid19Orch.GetCovid19Data(request)

	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}
