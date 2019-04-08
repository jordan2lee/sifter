package manager

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/sifter/emitter"
)

type Runtime struct {
	man         *Manager
	output      emitter.Emitter
	dir         string
	name        string
	Status      string
	VertexCount int64
	EdgeCount   int64
}

func (run *Runtime) NewTask(inputs map[string]interface{}) *Task {
	return &Task{run.man, run, run.dir, inputs}
}

func (run *Runtime) Close() {
	if run.output != nil {
		run.output.Close()
	}
	run.man.DropRuntime(run.name)
}

func (run *Runtime) EmitVertex(v *gripql.Vertex) error {
	atomic.AddInt64(&run.VertexCount, 1)
	return run.output.EmitVertex(v)
}

func (run *Runtime) EmitEdge(e *gripql.Edge) error {
	atomic.AddInt64(&run.EdgeCount, 1)
	return run.output.EmitEdge(e)
}

func (m *Runtime) Printf(s string, x ...interface{}) {
	c := fmt.Sprintf(s, x...)
	log.Printf(c)
	m.Status = c
}

func (m *Runtime) GetCurrent() string {
	return m.Status
}

func (m *Runtime) GetVertexCount() int64 {
	return m.VertexCount
}

func (m *Runtime) GetEdgeCount() int64 {
	return m.EdgeCount
}

func (m *Runtime) GetStepNum() int64 {
	return 1
}

func (m *Runtime) GetStepTotal() int64 {
	return 10
}
