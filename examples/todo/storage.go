package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TODOStorage struct {
	path string
}

func NewTODOStorage(path string) *TODOStorage {
	return &TODOStorage{path: path}
}

func (this *TODOStorage) Read(queries ...interface{}) {
	for _, rawQuery := range queries {
		query, ok := rawQuery.(*SelectTODOs)
		if ok {
			this.load(query)
		}
	}
}

func (this *TODOStorage) load(query *SelectTODOs) {
	raw, err := ioutil.ReadFile(this.path)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(raw, &query.Results)
	if err != nil {
		log.Println(err)
	}
}

func (this *TODOStorage) Write(instructions ...interface{}) {
	for _, rawCommand := range instructions {
		switch command := rawCommand.(type) {
		case InsertTODO:
			this.insert(command)
		case UpdateTODO:
			this.update(command)
		}
	}
}

func (this *TODOStorage) update(command UpdateTODO) {
	query := &SelectTODOs{}
	this.Read(query)
	for i, task := range query.Results {
		if task.Description == command.Description {
			query.Results[i].Completed = true
		}
	}
	this.store(query.Results)
}

func (this *TODOStorage) insert(command InsertTODO) {
	query := &SelectTODOs{}
	this.Read(query)

	query.Results = append(query.Results, TODORecord{Description: command.Description})
	this.store(query.Results)
}

func (this *TODOStorage) store(records []TODORecord) {
	jsonBytes, err := json.MarshalIndent(records, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(this.path, jsonBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}
