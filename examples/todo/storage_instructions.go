package main

type LoadTODOsFromStorage struct {
	Results []StoredTODO
}

type InsertTODOIntoStorage struct {
	Description string
}

type UpdateTODOInStorage struct {
	Description string
}

type StoredTODO struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
