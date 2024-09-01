package main

import (
        "net/http"
        "time"
	"fmt"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
        go func() {
                for {
			fmt.Println("Inside for!")
                        opsProcessed.Inc()
                        time.Sleep(2 * time.Second)
                }
        }()
}

var (
        opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
                Name: "myapp_processed_ops_total",
                Help: "The total number of processed events",
        })
)

func main() {
	fmt.Println("Hello, World!")
        recordMetrics()

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":9080", nil)
}
