package perfmon

import (
	//	"encoding/json"
	//	"fmt"
	"log"
	"net/http"
	//	"time"

	"github.com/gorilla/websocket"
	"github.com/hinakaze/perfmon/data"
	"github.com/mkevac/debugcharts/bindata"
)

//func dataHandler(w http.ResponseWriter, r *http.Request) {
//	mutex.RLock()
//	defer mutex.RUnlock()

//	if e := r.ParseForm(); e != nil {
//		log.Print("error parsing form")
//		return
//	}

//	callback := r.FormValue("callback")

//	fmt.Fprintf(w, "%v(", callback)

//	w.Header().Set("Content-Type", "application/json")

//	encoder := json.NewEncoder(w)
//	encoder.Encode(data)

//	fmt.Fprint(w, ")")
//}

func handleAsset(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := bindata.Asset(path)
		if err != nil {
			log.Print(err)
			return
		}

		n, err := w.Write(data)
		if err != nil {
			log.Print(err)
			return
		}

		if n != len(data) {
			log.Print("wrote less than supposed to")
			return
		}
	}
}

func dataFeedHandler(w http.ResponseWriter, r *http.Request) {
	//	var (
	//		lastPing time.Time
	//		lastPong time.Time
	//	)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//	conn.SetPongHandler(func(s string) error {
	//		lastPong = time.Now()
	//		return nil
	//	})

	// read and discard all messages
	go func(c *websocket.Conn) {
		for {
			if _, _, err := c.NextReader(); err != nil {
				c.Close()
				break
			}
		}
	}(conn)

	data.StartWSData(conn)
	//	c := s.addConsumer()

	//	defer func() {
	//		s.removeConsumer(c.id)
	//		conn.Close()
	//	}()

	//	var i uint

	//	for u := range c.c {
	//		websocket.WriteJSON(conn, u)
	//		i += 1

	//		if i%10 == 0 {
	//			if diff := lastPing.Sub(lastPong); diff > time.Second*60 {
	//				return
	//			}
	//			now := time.Now()
	//			if err := conn.WriteControl(websocket.PingMessage, nil, now.Add(time.Second)); err != nil {
	//				return
	//			}
	//			lastPing = now
	//		}
	//	}
}
