package agent

import (
	"fmt"
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

var rt runtime.MemStats

type Gauge float64
type Counter int64

type Metric struct {
	Alloc         Gauge
	BuckHashSys   Gauge
	Frees         Gauge
	GCCPUFraction Gauge
	GCSys         Gauge
	HeapAlloc     Gauge
	HeapIdle      Gauge
	HeapInuse     Gauge
	HeapObjects   Gauge
	HeapReleased  Gauge
	HeapSys       Gauge
	LastGC        Gauge
	Lookups       Gauge
	MCacheInuse   Gauge
	MCacheSys     Gauge
	MSpanInuse    Gauge
	MSpanSys      Gauge
	Mallocs       Gauge
	NextGC        Gauge
	NumForcedGC   Gauge
	NumGC         Gauge
	OtherSys      Gauge
	PauseTotalNs  Gauge
	StackInuse    Gauge
	StackSys      Gauge
	Sys           Gauge
	TotalAlloc    Gauge

	RandomValue Gauge
	PollCount   Counter
}

func GetMetric(m *Metric, pollInterval int) {
	for {
		runtime.ReadMemStats(&rt)
		m.Alloc = Gauge(rt.Alloc)
		m.BuckHashSys = Gauge(rt.BuckHashSys)
		m.Frees = Gauge(rt.Frees)
		m.GCCPUFraction = Gauge(rt.GCCPUFraction)
		m.GCSys = Gauge(rt.GCSys)
		m.HeapAlloc = Gauge(rt.HeapAlloc)
		m.HeapIdle = Gauge(rt.HeapIdle)
		m.HeapInuse = Gauge(rt.HeapInuse)
		m.HeapObjects = Gauge(rt.HeapObjects)
		m.HeapReleased = Gauge(rt.HeapReleased)
		m.HeapSys = Gauge(rt.HeapSys)
		m.LastGC = Gauge(rt.LastGC)
		m.Lookups = Gauge(rt.Lookups)
		m.MCacheInuse = Gauge(rt.MCacheInuse)
		m.MCacheSys = Gauge(rt.MCacheSys)
		m.MSpanInuse = Gauge(rt.MSpanInuse)
		m.MSpanSys = Gauge(rt.MSpanSys)
		m.Mallocs = Gauge(rt.Mallocs)
		m.NextGC = Gauge(rt.NextGC)
		m.NumForcedGC = Gauge(rt.NumForcedGC)
		m.NumGC = Gauge(rt.NumGC)
		m.OtherSys = Gauge(rt.OtherSys)
		m.PauseTotalNs = Gauge(rt.PauseTotalNs)
		m.StackInuse = Gauge(rt.StackInuse)
		m.StackSys = Gauge(rt.StackSys)
		m.Sys = Gauge(rt.Sys)
		m.TotalAlloc = Gauge(rt.TotalAlloc)

		m.RandomValue = Gauge(rand.Float64())
		m.PollCount++

		//log.Println(m.PollCount)
		time.Sleep(time.Duration(pollInterval) * time.Second)
	}
}

func SendMetric(metric Metric, path string) {

	val := reflect.ValueOf(&metric)

	val = val.Elem()
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
		field := val.Field(fieldIndex)
		//fmt.Printf("\tField %v: type %v - val :%v\n", val.Type().Field(fieldIndex).Name, field.Type(), field)
		f := val.FieldByName(val.Type().Field(fieldIndex).Name)
		switch f.Kind() {
		case reflect.Int64:
			//log.Printf("%s/update/counter/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
			go func(fieldIndex int) {
				//log.Printf("%s/update/counter/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
				resp, err := handlers.AgentRequest(fmt.Sprintf("http://%s/update/counter/%s/%v", path, val.Type().Field(fieldIndex).Name, field))
				if err != nil {
					return
				}
				defer resp.Body.Close()

			}(fieldIndex)
		case reflect.Float64:
			//log.Printf("%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
			go func(fieldIndex int) {
				// log.Printf("%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
				resp, err := handlers.AgentRequest(fmt.Sprintf("http://%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field))
				if err != nil {
					return
				}
				defer resp.Body.Close()
				//_, err := http.Post(fmt.Sprintf("%s/update/gauge/%s/%v", targetUrl, val.Type().Field(fieldIndex).Name, field), "text/plain", nil)
				//if err != nil {
				//	return
				//}
			}(fieldIndex)
		}
	}
}
