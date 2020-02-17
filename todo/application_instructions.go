package todo

type TODO struct {
	ID          uint64
	Description string
}

type ListTODOs struct {
	Results []TODO
}

type AddTODO struct {
	ID          uint64
	Description string
}

type CompleteTODO struct {
	ID uint64
}
