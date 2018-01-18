package joyride

type Factory func(interface{}) Procedure

type Procedure interface {
	Read() []interface{}
	Execute()
	Write() []interface{}
	Dispatch() []interface{}
	Continue() Procedure
}

type (
	Reader interface {
		Read(...interface{})
	}
	Writer interface {
		Write(...interface{})
	}
	Dispatcher interface {
		Dispatch(...interface{})
	}
)
