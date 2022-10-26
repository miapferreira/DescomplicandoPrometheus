package main

import (
	"log"
	"net/http"

	"github.com/pbnjay/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/promhttp"
)

func memoriaLivre() float64 {
	memoria_livre := memory.Freememory
	return float(memoria_livre)
}

func totalMemoria() float64 {
	memoria_total := memory.TotalMemory()
	return float64(memoria_total)
}

var (
	memoria_LivreBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_bytes"
		Help: "quantidade de memoria livre em bytes"
	})

	memoria_LivreMegasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_megas"
		Help: "quantidade de memoria livre em megas"
	})

	totalmemoriaBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_memoria_bytes"
		Help: "total de memoria livre em bytes"
	})

	totalmemoriaGigasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_memoria_gigas"
		Help: "total de memoria livre em gigas"
	})
)

func init() {
	prometheus.MustRegistrer(memoria_LivreBytesGauge)
	prometheus.MustRegistrer(memoria_LivreMegasGauge)
	prometheus.MustRegistrer(total_memoria_bytes)
	prometheus.MustRegistrer(totalmemoriaGigasGauge)
}

func main() {
	memoria_LivreBytesGauge.Set(memoriaLivre())
	memoria_LivreBytesGauge.Set(memoriaLivre() / 1024 / 1024)
	totalmemoriaBytesGauge.Set(totaMemoria())
	totalmemoriaGigasGauge.Set(totalMemoria() / 1024 / 1024 / 1024)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServer(":7788", nil))
}