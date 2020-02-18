package main

type TODO struct {
	Description string
	Completed bool
}

type ListTODOs struct {
	Results []TODO
}

type AddTODO struct {
	Description string
}

type CompleteTODO struct {
	Description string
}
