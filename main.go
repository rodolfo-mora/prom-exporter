package main

import (
	"prom-exporter/pkg/api"
)

func main() {
	api := api.NewAPI()
	api.RegisterRoutes()
}
