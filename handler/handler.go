package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hinakaze/perfmon/data"
)

func DataFeedHandler(w http.ResponseWriter, r *http.Request) {
	//	var (
	//		lastPing time.Time
	//		lastPong time.Time
	//	)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

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

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(index))
}

var index string = `
<!DOCTYPE html>
<html>
	<head>
		<title></title>
		<meta charset="utf-8" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.17.1/moment.min.js"></script>
		<script src="https://cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></script>
        <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
		<style type="text/css">
			.chart{
				min-width: 300px;
				width: 48%;
				height: 400px;
				margin-left: auto;
				margin-right: auto;
				display: inline-block;
			}
			h1{
				text-align: center;
			}
		</style>
	</head>
	<body>
		<h1>Sys Info</h1>
		<div id="c_mem" class="chart"></div>
		<div id="c_gc" class="chart"></div>
		<div id="c_goroutine" class="chart"></div>
		<div id="c_thread" class="chart"></div>
		
		<h1>Target Info</h1>
		<div id="c_latency_histo" class="chart"></div>
		<div id="c_latency_pie" class="chart"></div>
		
		<script src="http://localhost:8080/static/main.js"></script>
	</body>
</html>


`
