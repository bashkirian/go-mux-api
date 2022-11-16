package main

import "os"

// @title Balance microservice
// @version 1.0
// @description Balance service task made for Avito internship

// @host localhost:8010 for local, balance-db:8010 for docker
// @BasePath /
func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
