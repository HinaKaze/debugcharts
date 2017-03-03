package data

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WSData struct {
	Type  int
	Name  string
	Value float64
	Time  int64
}

var wsDataChan chan WSData
var RunFlag bool = false

func StartWSData(ws *websocket.Conn) {
	wsDataChan = make(chan WSData, 16)
	go func() {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("WSData server occur a error [%s],stopped it", x)
				close(wsDataChan)
				RunFlag = false
			}
		}()
		for {
			wsData := <-wsDataChan
			err := ws.WriteJSON(wsData)
			if err != nil {
				panic(err.Error())
			}
		}
	}()
	RunFlag = true
}

func SendWSData(name string, value float64) {
	go func() {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("SendWSData server occur a error [%s],stopped it", x)
				RunFlag = false
			}
		}()

		d := WSData{Type: 1, Name: name, Value: value, Time: time.Now().Unix() * 1000}
		wsDataChan <- d
	}()
}
