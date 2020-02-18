package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"runtime/debug"
)

type TODORecord struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type SelectTODOs struct {
	Results []TODORecord
}

type InsertTODO struct {
	Description string
}

type UpdateTODO struct {
	Description string
}

type TODOStorage struct {
	path string
}

func NewTODOStorage(path string) *TODOStorage {
	return &TODOStorage{path: path}
}

func (this *TODOStorage) Read(queries ...interface{}) {
	log.Println("Read: ", queries)
	debug.PrintStack()
	for _, rawQuery := range queries {
		query, ok := rawQuery.(*SelectTODOs)
		if !ok {
			continue
		}
		raw, err := ioutil.ReadFile(this.path)
		if err != nil {
			log.Println(err)
			continue
		}
		err = json.Unmarshal(raw, &query.Results)
		if err != nil {
			log.Println(err)
		}
	}
}

func (this *TODOStorage) Write(instructions ...interface{}) {
	log.Println("Write: ", instructions)
	for _, rawCommand := range instructions {
		switch command := rawCommand.(type) {
		case InsertTODO:
			query := &SelectTODOs{}
			this.Read(query)
			query.Results = append(query.Results, TODORecord{Description:command.Description})
			jsonBytes, err := json.Marshal(query.Results) //TODO refactor out duplication
			if err != nil {
				log.Println(err)
				continue
			}
			err = ioutil.WriteFile(this.path, jsonBytes, 0644)
			if err != nil {
				log.Println(err)
			}
		//case UpdateTODO:
		//	query := &SelectTODOs{}
		//	this.Read(query)
		//	for i, task := range query.Results {
		//		if task.Description == command.Description {
		//			query.Results[i].Completed = true
		//		}
		//	}
		//	jsonBytes, err := json.Marshal(query.Results) //TODO refactor out duplication
		//	if err != nil {
		//		log.Println(err)
		//		continue
		//	}
		//	err = ioutil.WriteFile(this.path, jsonBytes, 0644)
		//	if err != nil {
		//		log.Println(err)
		//	}
		}
	}
}
