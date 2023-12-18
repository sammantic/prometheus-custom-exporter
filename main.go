package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	promRegistry := prometheus.NewRegistry() // Init new registry

	// Create a gauge metric
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature_kelvin",
			Help: "Temperature in kelvin",
		},
		[]string{"location"},
	)

	promRegistry.Register(temp) // Register temp gauge metric

	temp.WithLabelValues("outside").Set(273.14)
	temp.WithLabelValues("inside").Set(298.44)

	fmt.Println("Exporter server: http://localhost:8000/metrics") // Print server info

	handler := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)
	http.ListenAndServe(":8000", nil)
}
