package main

import (
	"log"
	"net/http"

	"github.com/Haameed/bigip_exporter/internal/config"
	bigIPHTTP "github.com/Haameed/bigip_exporter/pkg/http"
	"github.com/Haameed/bigip_exporter/pkg/probe"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("Initialization error: %+v", err)
	}

	savedConfig := config.GetConfig()
	if err := bigIPHTTP.Configure(savedConfig); err != nil {
		log.Fatalf("%+v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/probe", probe.Handler)
	go func() {
		if err := http.ListenAndServe(savedConfig.Listen, nil); err != nil {
			log.Fatalf("Unable to serve: %v", err)
		}
	}()
	log.Printf("F5 bigips exporter is running, And listening on %q", savedConfig.Listen)
	select {}
}
