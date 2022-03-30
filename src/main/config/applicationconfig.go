package config

import (
	"github.com/covid19-service/src/main/apis"
	"github.com/covid19-service/src/main/client"
	"github.com/covid19-service/src/main/dal/repos"
	"github.com/covid19-service/src/main/service/impl"
	"github.com/covid19-service/src/main/service/orch"
	"github.com/labstack/echo"
)

func InitializeApplicationConfig(route *echo.Group) {
	googleClient := client.NewGoogleClient()
	dbConnections := GetDatabase()
	redisConnection := GetRedisClient()
	covid19Repo := repos.NewCovid19Repo(dbConnections)
	covid19Service := impl.NewCovid19Impl(covid19Repo)
	covid19Orch := orch.NewCovid19Orch(redisConnection, googleClient, covid19Service)
	apis.NewCovid19Controller(route, covid19Orch)
}
