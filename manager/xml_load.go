package manager

import (
	"fmt"
	"log"
  "io"
	"os"
	"sync"
	"github.com/bmeg/sifter/evaluate"
  "bufio"
  "encoding/xml"
)

type XMLLoadStep struct {
	Input         string        `json:"input"`
	Transform     TransformPipe `json:"transform"`
	SkipIfMissing bool          `json:"skipIfMissing"`
}

func xmlStream(file io.Reader, out chan map[string]interface{}) {
	buffer := bufio.NewReaderSize(file, 1024*1024*256) // 33554432
	decoder := xml.NewDecoder(buffer)

	nameStack := []string{}
	mapStack := []map[string]interface{}{}
	curString := []byte{}
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			nameStack = append(nameStack, se.Name.Local)
			mapStack = append(mapStack, map[string]interface{}{})
			curString = []byte{}
		case xml.EndElement:
			cMap := mapStack[len(mapStack)-1]
			nameStack = nameStack[0 : len(nameStack)-1]
			mapStack = mapStack[0 : len(mapStack)-1]
			if len(cMap) > 0 {
				//cMap["__text__"] = string(curString)
				mapStack[len(mapStack)-1][se.Name.Local] = cMap
			} else {
				if a, ok := mapStack[len(mapStack)-1][se.Name.Local]; ok {
					if aa, ok := a.([]string); ok {
						aa = append(aa, string(curString))
						mapStack[len(mapStack)-1][se.Name.Local] = aa
					} else {
						if as, ok := a.(string); ok {
							aa := []string{as, string(curString)}
							mapStack[len(mapStack)-1][se.Name.Local] = aa
						} else {
							log.Printf("Typing Error")
						}
					}
				} else {
					mapStack[len(mapStack)-1][se.Name.Local] = string(curString)
				}
			}
		case xml.CharData:
			curString = append(curString, se...)
		default:
			log.Printf("%#v\n", se)
		}
		if len(nameStack) == 1 {
			c := mapStack[0]
			if len(c) > 0 {
        out <- c
			}
		}
	}
}

func (ml *XMLLoadStep) Run(task *Task) error {
	log.Printf("Starting XML Load")
	input, err := evaluate.ExpressionString(ml.Input, task.Inputs, nil)
	inputPath, err := task.Path(input)
	if err != nil {
		return err
	}

	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		if ml.SkipIfMissing {
			return nil
		}
		return fmt.Errorf("File Not Found: %s", input)
	}
	log.Printf("Loading: %s", inputPath)

  file, err := os.Open(inputPath)
  if err != nil {
    return err
  }

	wg := &sync.WaitGroup{}
	procChan := make(chan map[string]interface{}, 100)

	ml.Transform.Start(procChan, task, wg)
  xmlStream(file, procChan)
  log.Printf("Done Loading")
	close(procChan)
	wg.Wait()

	return nil
}
