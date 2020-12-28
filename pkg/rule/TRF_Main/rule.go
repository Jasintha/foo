package TRF_Main

import (
	"foo/pkg/functions"
)

type CFGContext struct {
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext() *CFGContext {
	return &CFGContext{}
}
func TRF_Main() interface{} {

	cfg := NewCFGContext()
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiHybridFunctionNodeM0()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM0() error {
	var err error
	functions.HF_Main()
	return err
}
