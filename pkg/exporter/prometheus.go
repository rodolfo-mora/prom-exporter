package exporter

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Name = "hostcheck"
	Help = "Gauge values meaning. Host status - up: 1, down: 0"
)

type Prometheus struct {
	Gauge   *prometheus.GaugeVec
	Hosts   chan string
	Tracker []string
}

func NewPrometheusExporter() Prometheus {
	// var t tracker.Tracker
	hosts := make(chan string, 1)
	prom := Prometheus{
		Gauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: Name,
				Help: Help,
			},
			[]string{"hostname"},
		),
		Tracker: []string{},
		Hosts:   hosts,
	}
	prometheus.MustRegister(prom.Gauge)
	return prom
}

func (p *Prometheus) Register(hostname string) {
	p.Gauge.WithLabelValues(hostname).Set(float64(1))
}

func (p Prometheus) Export() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func (p *Prometheus) Track(name string) {
	p.Tracker = append(p.Tracker, name)
}

func (p Prometheus) RandomHostDown() {
	for range time.Tick(60 * time.Second) {
		p.HostDown("ninja.test.123")
	}
}

func (p Prometheus) HostDown(hostname string) {
	p.Gauge.WithLabelValues(hostname).Set(float64(0))
}

func (p Prometheus) Display() {
	log.Println(p.Tracker)
}
