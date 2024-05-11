package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
)

var (
	HandlerRPS = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "handler_rps_total",
			Help: "The total number of processed requests with response codes",
		},
		[]string{"handler", "code"},
	)

	HandlerRT = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "handler_rt",
			Help: "Responce time of requests",
		},
		[]string{"handler"},
	)
)

func Register() {
	err := prometheus.Register(HandlerRT)
	if err != nil {
		log.Printf("error registering handler: %v", err)
	}

	err = prometheus.Register(HandlerRPS)
	if err != nil {
		log.Printf("error registering handler: %v", err)
	}
}
