package db

import (
	"log"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
}

func TestQuery(t *testing.T) {
	db := New()
	metricName := "http_request"
	labels := []Label{
		{
			K: "ip",
			V: "192.168.2.11",
		},
	}
	values := []Sample{
		{
			Timestamp: time.Now().Unix(),
			Value:     "1.1",
		},
	}
	db.Save(metricName, labels, values)

	r := db.Query(metricName, labels)

	if len(r) != len(values) {
		log.Fatalf("expected len=%d, got len=%d", len(values), len(r))
	}

	for i, e := range r {
		expectedElement := values[i]
		actualElement := e

		if expectedElement.Timestamp != actualElement.Timestamp {
			log.Fatalf("expected timestamp=%d, got timestamp=%d", expectedElement.Timestamp, actualElement.Timestamp)
		}
		if expectedElement.Value != actualElement.Value {
			log.Fatalf("expected value=%d, got value=%d", expectedElement.Value, actualElement.Value)
		}
	}

}