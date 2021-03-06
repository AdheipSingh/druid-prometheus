package main

import (
	"net/http"
	"time"

	"druid-prometheus/collector"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.

	uptimeTicker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			go executeCronJob()
			time.Sleep(500 * time.Millisecond)
			http.Handle("/metrics", promhttp.Handler())
			log.Printf("Beginning to serve on port :8080")
			log.Fatal(http.ListenAndServe(":8080", nil))
		}
	}

}

func executeCronJob() {
	metric := collector.Collector()
	prometheus.MustRegister(metric)
}
