package joyride

import "context"

type nopIO struct{}

func (nopIO) Dispatch(_ context.Context, _ ...interface{}) {}
func (nopIO) Read(_ context.Context, _ ...interface{})     {}
func (nopIO) Write(_ context.Context, _ ...interface{})    {}
