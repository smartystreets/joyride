package main

import (
	"context"
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

func (this *TODOStorage) Read(_ context.Context, queries ...interface{}) {
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

func (this *TODOStorage) Write(ctx context.Context, instructions ...interface{}) {
	for _, rawCommand := range instructions {
		switch command := rawCommand.(type) {
		case InsertTODOIntoStorage:
			this.insert(ctx, command)
		case UpdateTODOInStorage:
			this.update(ctx, command)
		}
	}
}

func (this *TODOStorage) update(ctx context.Context, command UpdateTODOInStorage) {
	query := &LoadTODOsFromStorage{}
	this.Read(ctx, query)
	for i, task := range query.Results {
		if task.Description == command.Description {
			query.Results[i].Completed = true
		}
	}
	this.store(query.Results)
}

func (this *TODOStorage) insert(ctx context.Context, command InsertTODOIntoStorage) {
	query := &LoadTODOsFromStorage{}
	this.Read(ctx, query)

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
