# covid19-service

To run the service : `go run main.go`

Web Framework:  : echo
DataBase : MongoDB
Cache : Redis

This service consists 2 Apis :

1. To Update the covid cases of india in a MongoDb

Endpoint : 127.0.0.1:8000/api/v1/covid/update
Method : POST


2. To get the total cases in a state via user's lat lng

EndPoint : 127.0.0.1:8000/api/v1/covid/cases
Method : GET
Request : {
    "lat": "31.2389",
    "lng": "76.0243",
    "api_key" : <api-key for reverse geo location>
}
  
  
