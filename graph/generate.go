package graph

import (
  "github.com/bmeg/sifter/schema"
)

func GenerateGraph(sc *schema.Schemas, class string, data map[string]interface{}, emitter GraphEmitter) error {
	if o, err := sc.Generate(class, data); err == nil {
		for _, j := range o {
			if j.Vertex != nil {
				emitter.EmitVertex(j.Vertex)
			} else if j.Edge != nil {
				emitter.EmitEdge(j.Edge)
			}
		}
	} else {
		return err
	}
	return nil
}
