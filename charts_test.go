package perfmon

import (
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go func() {
		for {
			<-time.After(time.Second)
			SendValue("latency", float64(rand.Intn(100)))
		}
	}()
	<-time.After(time.Hour)
}
