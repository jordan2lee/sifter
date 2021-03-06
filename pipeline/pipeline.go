package pipeline

import (
	"fmt"
	"log"
	"sync/atomic"

	//"github.com/bmeg/grip/gripql"
	"github.com/bmeg/sifter/datastore"
	"github.com/bmeg/sifter/emitter"
)

type Runtime struct {
	//man         *Manager
	output      emitter.Emitter
	dir         string
	name        string
	Status      string
	StepCount   int64
	StepTotal   int64
	OutputCallback func(string, string) error
	datastore   datastore.DataStore
}

func NewRuntime(output emitter.Emitter, dir string, name string, ds datastore.DataStore) *Runtime {
	return &Runtime{output:output, dir:dir, name:name, Status:"Starting", datastore:ds}
}

func (run *Runtime) NewTask(inputs map[string]interface{}) *Task {
	return &Task{Name: run.name, Runtime:run, Workdir:run.dir, Inputs:inputs, AllowLocalFiles:true, DataStore:run.datastore}
}

func (run *Runtime) Close() {
	if run.output != nil {
		run.output.Close()
	}
	//run.man.DropRuntime(run.name)
}

func (run *Runtime) EmitObject(prefix string, c string, o map[string]interface{}) error {
	if prefix == "" {
		return run.output.EmitObject(run.name, c,o)
	}
	return run.output.EmitObject(prefix, c,o)
}


func (run *Runtime) EmitTable(prefix string, columns []string) emitter.TableEmitter {
	return run.output.EmitTable(prefix, columns)
}

func (m *Runtime) Printf(s string, x ...interface{}) {
	c := fmt.Sprintf(s, x...)
	log.Printf(c)
	m.Status = c
}

func (m *Runtime) GetCurrent() string {
	return m.Status
}

func (m *Runtime) GetStepNum() int64 {
	return m.StepCount
}

func (m *Runtime) GetStepTotal() int64 {
	return m.StepTotal
}

func (m *Runtime) SetStepCountTotal(i int64) {
	m.StepTotal = i
}

func (m *Runtime) AddStepCount(i int64) {
	atomic.AddInt64(&m.StepCount, i)
}
