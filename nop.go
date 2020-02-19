package joyride

type nopIO struct{}

func (nopIO) Dispatch(...interface{}) {}
func (nopIO) Read(...interface{})     {}
func (nopIO) Write(...interface{})    {}
