# qpromethues

prometheus wrapper in golang only for so.com team

## Usage


``` golang
// you should init prometheus in your main.go first
func main() {
    // ...

    var Idc = "idc" // your cluster
    var ProjectName = "your_project_name" // your project name

    qprometheus.Init(qprometheus.Opts{
        Idc:             Idc,
        AppName:         ProjectName,
        
        // the prometheus buckets 
        HistogramBucket: []float64{1, 5, 10, 20, 30, 35, 40, 50, 60, 76, 80, 90, 100, 200, 400, 600, 1000, 1200, 1300, 1500, 1600, 1800, 2000, 2300, 2500, 2700, 3000, 3300, 3500, 3800, 4000, 4300, 4500, 4800, 5000, 5300, 5500, 5800, 6000, 6300, 6500, 6800, 7000, 7300, 7500, 7800, 8000, 9000, 10000, 11000},

        // path in WatchPath will be automatically recorded only if you have a middleware 
        WatchPath: map[string]struct{}{
            "/json":       {},
            "/_stats":     {},
            "/v1/example": {},
        },
    })

    // path: expose /metrics for prometheus 
    qprometheus.MetricsServerStart("/metrics", 18081)

    // ...
}


// record the qps counter
qprometheus.GetWrapper().QpsCountLog(qprometheus.QPSRecord{
    Times:  1,                       // the count you want to record
    Api:    "/your-custom-api-name", 
    Module: "your-custom-module",    
    Method: "GET",                   // the request method
    Code:   200,                     // the request code
})

// record the latency
qprometheus.GetWrapper().LatencyLog(qprometheus.LatencyRecord{
    Time:   1000, // time in millisecond
    Api:    "/your-custom-api-name",
    Module: "your-custom-module",
    Method: "GET",
})
    
```
