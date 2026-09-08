package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	},
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()

	log.Printf("request received: %s %s", r.Method, r.URL.Path)

	fmt.Fprintln(w, "Hello from Go!")
}

func main() {
	prometheus.MustRegister(httpRequests)

	http.HandleFunc("/", helloHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("server starting on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
