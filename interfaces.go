package joyride

type Factory func(interface{}) Procedure

type Procedure interface {
	Populate() []interface{}
	Process()
	Persist() []interface{}
	Publish() []interface{}
	Procedure() Procedure
}

type (
	Reader interface {
		Read(...interface{})
	}
	Writer interface {
		Write(...interface{})
	}
	Publisher interface {
		Publish(...interface{})
	}
)
