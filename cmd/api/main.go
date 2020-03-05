package main

import (
	"github.com/armando0194/movie-night-backend/pkg/api"
	"github.com/armando0194/movie-night-backend/pkg/utl/config"
)

func main() {

	// TODO read enviroment from env variables

	cfg, err := config.Load("dev")
	checkErr(err)

	checkErr(api.Serve(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
