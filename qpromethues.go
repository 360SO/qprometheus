package qprometheus

import (
	"fmt"
	"strings"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 选项配置
type prom struct {
	Appname   string
	Idc       string
	WatchPath map[string]struct{}
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

var Wrapper *prom

// 选项配置
type Opts struct {
	AppName         string              // 项目名称
	Idc             string              // 机房名称
	WatchPath       map[string]struct{} // 监控路径
	HistogramBucket []float64           // 桶距配置
}

// 初始化
func Init(opts Opts) {
	if strings.TrimSpace(opts.AppName) == "" {
		panic("Prometheus Opts.AppName Can't Be Empty")
	}

	if strings.TrimSpace(opts.Idc) == "" {
		panic("Prometheus Opts.Idc Can't Be Empty")
	}

	if len(opts.HistogramBucket) == 0 {
		panic("Prometheus Opts.HistogramBucket Can't Be Empty")
	}

	p := &prom{
		Appname:   opts.AppName,
		Idc:       opts.Idc,
		WatchPath: opts.WatchPath,
		counter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "module_responses",
				Help: "used to calculate qps, failure ratio",
			},
			[]string{"app", "module", "api", "method", "code", "idc"},
		),
		histogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "response_duration_milliseconds",
				Help:    "HTTP latency distributions",
				Buckets: opts.HistogramBucket,
			},
			[]string{"app", "module", "api", "method", "idc"},
		),
	}

	prometheus.MustRegister(p.counter)
	prometheus.MustRegister(p.histogram)

	Wrapper = p
}

// 启动监听server，收集metrics数据
func MetricsServerStart(path string, port int) {
	// prometheus metrics path
	go func() {
		http.Handle(path, promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		fmt.Printf("Prometheus start with path '/metrics' and port on %d\n", port)
	}()
}
