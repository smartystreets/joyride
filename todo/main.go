package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/smartystreets/joyride"
)

func main() {
	log.SetFlags(log.Lshortfile)
	var description string
	flag.StringVar(&description, "desc", "", "description of the task")
	flag.Parse()
	storage := NewTODOStorage("todo.json")
	runner := joyride.NewRunner(storage, storage, NopDispatcher{})
	handler := NewHandler(runner)
	if description != "" {
		handler.Handle(AddTODO{Description: description})
	}

	instruction := &ListTODOs{}
	handler.Handle(instruction)
	for i, result := range instruction.Results {
		fmt.Printf("%d. [%s] %s\n", i+1, completion[result.Completed], result.Description)
	}

}

var completion = map[bool]string{
	false: " ",
	true:  "x",
}

type NopDispatcher struct{}

func (n NopDispatcher) Dispatch(...interface{}) {}
