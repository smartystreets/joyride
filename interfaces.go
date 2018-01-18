package joyride

type Factory func(interface{}) Procedure

type Procedure interface {
	Reads() []interface{}
	Execute()
	Writes() []interface{}
	Messages() []interface{}
	Next() Procedure
	Finally()
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
