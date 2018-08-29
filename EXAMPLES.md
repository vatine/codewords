Here's the "how to create metrics" examples from the slide deck:


```go
// creating a counter metric
myCnt := prometheus.NewCounter(prometheus.CounterOpts{Name: "…", Help: "…",})
// registering
prometheus.MustRegister(myCnt)
// Incrementing
myCnt.Inc()
```

```go
// creating a histogram metric
myHisto := prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "…", Help: "…",
		Buckets: prometheus.LinearBuckets(1, 3, 5),
	}
)
// registering
prometheus.MustRegister(myHisto)
// using
myHisto.Observe(3.14159)
```

```go
// creating a gauge metric
myGauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: "…", Help: "…",})
// registering
prometheus.MustRegister(myGauge)
// using
my.Gauge.Set(2.718281)
```
