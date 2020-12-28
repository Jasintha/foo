package ARF_GetAllPersons

import (
	"foo/pkg/env"
	"foo/pkg/model"
)

type CFGContext struct {
	_xiDBNode_UserName interface{}
	_xiDBNode_Password interface{}
	allPersons         []*model.Person
	_returnValue       interface{}
	_errorValue        interface{}
}

func NewCFGContext() *CFGContext {
	return &CFGContext{}
}
func ARF_GetAllPersons() interface{} {

	cfg := NewCFGContext()
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiDBNodeM01()
	return nil
}

func (cfg *CFGContext) xiDBNodeM01() error {
     person := model.Person{}
	db := person.PreloadPerson(env.RDB)
	_target := []*model.Person{}
	db.Find(&_target)
	cfg.allPersons = _target
	cfg._returnValue = _target
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
