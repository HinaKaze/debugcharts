package data

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/shirou/gopsutil/process"
)

func GatherData(send func(string, float64)) {
	myProcess, _ := process.NewProcess(int32(os.Getpid()))
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
				send("gc", float64(ms.PauseNs[(gcNum+255)%256]))
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
				send("thread", float64(threadsNum))
			}
		}
	}()

	go func() {
		timer := time.Tick(time.Second)
		for {
			<-timer
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			send("mem", float64(ms.Alloc))
			send("goroutine", float64(pprof.Lookup("goroutine").Count()))
		}
	}()

}
