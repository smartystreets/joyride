package joyride

type (
	Factory   func(interface{}) Procedure
	Procedure interface {
		Read() []interface{}
		Execute()
		Write() []interface{}
		Dispatch() []interface{}
		Continue() Procedure
	}
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
