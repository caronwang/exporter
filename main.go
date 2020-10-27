package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

func newFooCollector() *fooCollector {
	m1 := make(map[string]string)
	m1["env"] = "prod"
	v := []string{"hostname"}
	return &fooCollector{
		fooMetric: prometheus.NewDesc("fff_metrics", "Show metrics a for mysql", nil, nil),
		barMetric: prometheus.NewDesc("bbb_metrics", "Show metrics a bar occu", v, m1),
	}
}

func (collect *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collect.barMetric
	ch <- collect.fooMetric

}

var cnt = 0

func (collect *fooCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	fmt.Println("collect...")
	fmt.Println("cnt:", cnt)
	cnt++
	ch <- prometheus.MustNewConstMetric(collect.fooMetric, prometheus.GaugeValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collect.barMetric, prometheus.CounterValue, float64(cnt), "total")
}

func main() {
	var port string
	flag.StringVar(&port, "p", "8000", "-p port")
	flag.Parse()
	foo := newFooCollector()
	prometheus.MustRegister(foo)
	log.Info("beging to server on Port: " + port)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
