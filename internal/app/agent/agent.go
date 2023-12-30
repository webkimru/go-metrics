package agent

import (
	"fmt"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent/metrics"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

var rt runtime.MemStats

func GetMetric(m *metrics.Metric, pollInterval int) {
	pollDuration := time.Duration(pollInterval) * time.Second

	for {
		runtime.ReadMemStats(&rt)
		m.Alloc = metrics.Gauge(rt.Alloc)
		m.BuckHashSys = metrics.Gauge(rt.BuckHashSys)
		m.Frees = metrics.Gauge(rt.Frees)
		m.GCCPUFraction = metrics.Gauge(rt.GCCPUFraction)
		m.GCSys = metrics.Gauge(rt.GCSys)
		m.HeapAlloc = metrics.Gauge(rt.HeapAlloc)
		m.HeapIdle = metrics.Gauge(rt.HeapIdle)
		m.HeapInuse = metrics.Gauge(rt.HeapInuse)
		m.HeapObjects = metrics.Gauge(rt.HeapObjects)
		m.HeapReleased = metrics.Gauge(rt.HeapReleased)
		m.HeapSys = metrics.Gauge(rt.HeapSys)
		m.LastGC = metrics.Gauge(rt.LastGC)
		m.Lookups = metrics.Gauge(rt.Lookups)
		m.MCacheInuse = metrics.Gauge(rt.MCacheInuse)
		m.MCacheSys = metrics.Gauge(rt.MCacheSys)
		m.MSpanInuse = metrics.Gauge(rt.MSpanInuse)
		m.MSpanSys = metrics.Gauge(rt.MSpanSys)
		m.Mallocs = metrics.Gauge(rt.Mallocs)
		m.NextGC = metrics.Gauge(rt.NextGC)
		m.NumForcedGC = metrics.Gauge(rt.NumForcedGC)
		m.NumGC = metrics.Gauge(rt.NumGC)
		m.OtherSys = metrics.Gauge(rt.OtherSys)
		m.PauseTotalNs = metrics.Gauge(rt.PauseTotalNs)
		m.StackInuse = metrics.Gauge(rt.StackInuse)
		m.StackSys = metrics.Gauge(rt.StackSys)
		m.Sys = metrics.Gauge(rt.Sys)
		m.TotalAlloc = metrics.Gauge(rt.TotalAlloc)

		m.RandomValue = metrics.Gauge(rand.Float64())
		m.PollCount++

		//log.Println(m.PollCount)
		time.Sleep(pollDuration)
	}
}

func SendMetric(metric metrics.Metric, path string) {

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
				err := Send(fmt.Sprintf("http://%s/update/counter/%s/%v", path, val.Type().Field(fieldIndex).Name, field))
				if err != nil {
					log.Println(err)
				}

			}(fieldIndex)
		case reflect.Float64:
			//log.Printf("%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
			go func(fieldIndex int) {
				// log.Printf("%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field)
				err := Send(fmt.Sprintf("http://%s/update/gauge/%s/%v", path, val.Type().Field(fieldIndex).Name, field))
				if err != nil {
					log.Println(err)
				}
				//_, err := http.Post(fmt.Sprintf("%s/update/gauge/%s/%v", targetUrl, val.Type().Field(fieldIndex).Name, field), "text/plain", nil)
				//if err != nil {
				//	return
				//}
			}(fieldIndex)
		}
	}
}

func Send(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	return nil
}
