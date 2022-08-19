package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"prom-exporter/pkg/exporter"
	"prom-exporter/pkg/persister"

	"github.com/gorilla/mux"
)

type Systole struct {
	Name string `json:"name"`
}

type API struct {
	exporter  exporter.Prometheus
	router    *mux.Router
	persister persister.Persister
}

func NewAPI(r *mux.Router) API {
	ctx := context.TODO()
	cli := persister.NewRedisClient()
	a := API{
		exporter:  exporter.NewPrometheusExporter(),
		router:    r,
		persister: persister.NewRedisPersister(ctx, cli),
	}

	go a.exporter.Export()
	go a.exporter.RandomHostDown()
	return a
}

func (a *API) RegisterRoutes() {
	log.Println("Launching API")
	a.router.HandleFunc("/api/v1/hostcheck", a.Hostcheck).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", a.router))
}

func (a *API) Hostcheck(w http.ResponseWriter, r *http.Request) {
	var beat Systole
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)

	err = json.Unmarshal(body, &beat)
	handleError(err)

	a.exporter.Register(beat.Name)
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
