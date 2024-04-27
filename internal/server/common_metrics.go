package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
)

var (
	handlerRPS = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "handler_rps_total",
			Help: "The total number of processed requests with response codes",
		},
		[]string{"handler", "code"},
	)

	handlerRT = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "handler_rt",
			Help: "Responce time of requests",
		},
		[]string{"handler"},
	)
)

func Register() {
	err := prometheus.Register(handlerRT)
	if err != nil {
		log.Printf("error registering handler: %v", err)
	}

	err = prometheus.Register(handlerRPS)
	if err != nil {
		log.Printf("error registering handler: %v", err)
	}
}
