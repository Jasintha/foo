package api

import (
	arfaddperson "foo/pkg/rule/ARF_AddPerson"
	"reflect"
)

import (
	"foo/pkg/model"
	"foo/pub"
	"github.com/jinzhu/copier"
)

func AddPerson(person *pub.Person) (*pub.Person, error) {

	mPerson := model.Person{}
	copier.Copy(&mPerson, &person)
	result := arfaddperson.ARF_AddPerson(&mPerson)
	if reflect.TypeOf(result) == reflect.TypeOf(pub.Person{}) || reflect.TypeOf(result) == reflect.TypeOf(&pub.Person{}) {
		return nil, nil
	} else {
		return nil, nil
	}
}
