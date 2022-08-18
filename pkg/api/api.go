package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"prom-exporter/pkg/exporter"

	"github.com/gorilla/mux"
)

type Systole struct {
	Name string `json:"name"`
}

type API struct {
	Exporter exporter.Prometheus
}

func NewAPI() API {
	a := API{
		Exporter: exporter.NewPrometheusExporter(),
	}

	go a.Exporter.Export()
	go a.Exporter.RandomHostDown()
	return a
}

func (a *API) RegisterRoutes() {
	log.Println("Launching API")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/hostcheck", a.Beat).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (a *API) Beat(w http.ResponseWriter, r *http.Request) {
	var beat Systole
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)

	err = json.Unmarshal(body, &beat)
	handleError(err)

	a.Exporter.Register(beat.Name)
	// a.Exporter.Track(beat.Name)
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
