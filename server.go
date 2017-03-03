package perfmon

import (
	"runtime"
	"runtime/pprof"
	"time"

	//	"github.com/gorilla/websocket"

	//	datas "github.com/hinakaze/debugcharts/data"
)

//type server struct {
//	consumers      []consumer
//	consumersMutex sync.RWMutex
//}

//type consumer struct {
//	id uint
//	c  chan update
//}

//type update struct {
//	Ts             int64  //ms now
//	BytesAllocated uint64 //mem allo
//	GcPause        uint64
//	CpuUser        float64
//	CpuSys         float64
//}

//func (s *server) sendToConsumers(u update) {
//	s.consumersMutex.RLock()
//	defer s.consumersMutex.RUnlock()

//	for _, c := range s.consumers {
//		c.c <- u
//	}
//}

//func (s *server) removeConsumer(id uint) {
//	s.consumersMutex.Lock()
//	defer s.consumersMutex.Unlock()

//	var consumerId uint
//	var consumerFound bool

//	for i, c := range s.consumers {
//		if c.id == id {
//			consumerFound = true
//			consumerId = uint(i)
//			break
//		}
//	}

//	if consumerFound {
//		s.consumers = append(s.consumers[:consumerId], s.consumers[consumerId+1:]...)
//	}
//}

//func (s *server) addConsumer() consumer {
//	s.consumersMutex.Lock()
//	defer s.consumersMutex.Unlock()

//	lastConsumerId += 1

//	c := consumer{
//		id: lastConsumerId,
//		c:  make(chan update),
//	}

//	s.consumers = append(s.consumers, c)

//	return c
//}

func gatherData() {
	//gc
	go func() {
		var gcNum uint32 = 0
		timer := time.Tick(time.Second)
		for {
			<-timer
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			nowNum := ms.NumGC
			if gcNum < nowNum {
				gcNum = nowNum
				SendValue("gc", float64(ms.PauseNs[(gcNum+255)%256]))
			}
		}
	}()

	//thread
	go func() {
		timer := time.Tick(time.Second * 5)
		for {
			<-timer
			threadsNum, err := myProcess.NumThreads()
			if err == nil {
				SendValue("thread", float64(threadsNum))
			}
		}
	}()

	go func() {
		timer := time.Tick(time.Second)
		for {
			<-timer
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			SendValue("mem", float64(ms.Alloc))
			SendValue("goroutine", float64(pprof.Lookup("goroutine").Count()))
		}
	}()

}
