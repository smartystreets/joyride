package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/smartystreets/joyride/v2"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var description string
	var completed bool
	flag.StringVar(&description, "desc", "", "description of the task")
	flag.BoolVar(&completed, "complete", false, "completion status of the task")
	flag.Parse()

	storage := NewTODOStorage("todo.json")

	runner := joyride.NewRunner(joyride.WithStorageReader(storage), joyride.WithStorageWriter(storage))

	if description != "" && !completed {
		NewHandler(runner).Handle(AddTODO{Description: description})
	} else if description != "" && completed {
		NewHandler(runner).Handle(CompleteTODO{Description: description})
	}

	instruction := &ListTODOs{}
	NewHandler(runner).Handle(instruction)
	for i, result := range instruction.Results {
		fmt.Printf("%d. [%s] %s\n", i+1, completion[result.Completed], result.Description)
	}
}

var completion = map[bool]string{
	false: " ",
	true:  "x",
}
