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
	"log"
	"net/http"
	"os"

	datas "github.com/hinakaze/perfmon/data"
	"github.com/hinakaze/perfmon/handler"
)

var ()

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	log.Println(path)
	http.HandleFunc("/debug/charts/data-feed", handler.DataFeedHandler)
	http.HandleFunc("/debug/charts/", handler.Welcome)

	go datas.GatherData(datas.SendWSData)
}

func SendValue(name string, value float64) {
	if datas.RunFlag {
		datas.SendWSData(name, value)
	}
}
