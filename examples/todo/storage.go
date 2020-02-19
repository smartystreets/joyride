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
		query, ok := rawQuery.(*LoadTODOsFromStorage)
		if ok {
			this.load(query)
		}
	}
}

func (this *TODOStorage) load(query *LoadTODOsFromStorage) {
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
		case InsertTODOIntoStorage:
			this.insert(command)
		case UpdateTODOInStorage:
			this.update(command)
		}
	}
}

func (this *TODOStorage) update(command UpdateTODOInStorage) {
	query := &LoadTODOsFromStorage{}
	this.Read(query)
	for i, task := range query.Results {
		if task.Description == command.Description {
			query.Results[i].Completed = true
		}
	}
	this.store(query.Results)
}

func (this *TODOStorage) insert(command InsertTODOIntoStorage) {
	query := &LoadTODOsFromStorage{}
	this.Read(query)

	query.Results = append(query.Results, StoredTODO{Description: command.Description})
	this.store(query.Results)
}

func (this *TODOStorage) store(records []StoredTODO) {
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
