package main

type SelectTODOs struct {
	Results []TODORecord
}

type InsertTODO struct {
	Description string
}

type UpdateTODO struct {
	Description string
}

type TODORecord struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
