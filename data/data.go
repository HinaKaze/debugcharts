package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type Data struct {
	Time  int64 //data created time(ns)
	Value int
}

type Chart struct {
	Type     int //pre build charts type
	Name     string
	DataMode int // 0 active 1 passive
	Datas    []Data
	sync.RWMutex
}

func (this *Chart) Init(t int, name string, mode int) {
	this.Type = t
	this.Name = name
	this.DataMode = mode
	this.Datas = make([]Data, 0)
}

func (this *Chart) AddData(data Data) {
	this.Lock()
	defer this.Unlock()
	this.Datas = append(this.Datas, data)
}

func (this *Chart) Save() {
	this.RLock()
	defer this.RUnlock()
	fileName := fmt.Sprintf("%s-%d-%d-%d.json", this.Name, this.Type, this.DataMode, time.Now().Unix())
	bytes, err := json.Marshal(this.Datas)
	if err != nil {
		log.Printf("Chart save failed,error [%s]", err)
		return
	}
	err = ioutil.WriteFile(fileName, bytes, os.ModePerm)
	if err != nil {
		log.Printf("Chart save failed,error [%s]", err)
	}
	return
}

type Charts struct {
	Charts map[string]Chart
	sync.RWMutex
}

func (this *Charts) Init() {
	this.Charts = make(map[string]Chart)
}
