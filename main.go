package main

import (
	"prom-exporter/pkg/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := api.NewAPI(r)
	api.RegisterRoutes()
}
