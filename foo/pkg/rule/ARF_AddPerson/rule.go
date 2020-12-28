package ARF_AddPerson

import (
	"foo/pkg/model"
)

type CFGContext struct {
	Person             *model.Person
	newPerson          *model.Person
	_xiDBNode_UserName interface{}
	_xiDBNode_Password interface{}
	_returnValue       interface{}
	_errorValue        interface{}
}

func NewCFGContext(pPerson *model.Person) *CFGContext {
	return &CFGContext{

		Person: pPerson,
	}
}
func ARF_AddPerson(pPerson *model.Person) interface{} {

	cfg := NewCFGContext(pPerson)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiDBNodeM01()
	return nil
}

func (cfg *CFGContext) xiDBNodeM01() error {

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.newPerson = cfg.Person

	return nil
}
