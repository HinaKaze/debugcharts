// Simple live charts for memory consumption and GC pauses.
//
// To use debugcharts, link this package into your program:
//	import _ "github.com/mkevac/debugcharts"
//
// If your application is not already running an http server, you
// need to start one.  Add "net/http" and "log" to your imports and
// the following code to your main function:
//
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()
//
// Then go look at charts:
//
//	http://localhost:6060/debug/charts
//
package perfmon

import (
	"html/template"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	datas "github.com/hinakaze/perfmon/data"
	"github.com/shirou/gopsutil/process"
)

type SimplePair struct {
	Ts    uint64
	Value uint64
}

type CPUPair struct {
	Ts   uint64
	User float64
	Sys  float64
}

type DataStorage struct {
	BytesAllocated []SimplePair
	GcPauses       []SimplePair
	CpuUsage       []CPUPair
}

const (
	maxCount int = 86400
)

var (
	//	data           DataStorage
	lastPause      uint32
	mutex          sync.RWMutex
	lastConsumerId uint
	//	s              server
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	prevSysTime  float64
	prevUserTime float64
	myProcess    *process.Process
)

func init() {
	http.HandleFunc("/debug/charts/data-feed", dataFeedHandler)
	//	http.HandleFunc("/debug/charts/data", dataHandler)
	http.HandleFunc("/debug/charts/", Welcome)
	http.Handle("/debug/charts/static/", http.StripPrefix("/debug/charts/static/", http.FileServer(http.Dir("C:/develop/GOPATH/src/github.com/hinakaze/perfmon/static"))))
	//	http.HandleFunc("/debug/charts/main.js", handleAsset("static/main.js"))
	//	http.HandleFunc("/debug/charts/jquery-2.1.4.min.js", handleAsset("static/jquery-2.1.4.min.js"))
	//	http.HandleFunc("/debug/charts/moment.min.js", handleAsset("static/moment.min.js"))

	myProcess, _ = process.NewProcess(int32(os.Getpid()))

	// preallocate arrays in data, helps save on reallocations caused by append()
	// when maxCount is large
	//	data.BytesAllocated = make([]SimplePair, 0, maxCount)
	//	data.GcPauses = make([]SimplePair, 0, maxCount)
	//	data.CpuUsage = make([]CPUPair, 0, maxCount)

	go gatherData()
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("C:/develop/GOPATH/src/github.com/hinakaze/perfmon/static/index.html")
	t.Execute(w, nil)
}

func SendValue(name string, value float64) {
	if datas.RunFlag {
		datas.SendWSData(name, value)
	}
}
