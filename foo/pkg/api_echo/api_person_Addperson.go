package api_echo

import (
	model "foo/pkg/model"
	arfaddperson "foo/pkg/rule/ARF_AddPerson"
	"reflect"
)

import (
	"github.com/labstack/echo/v4"
)

func AddPerson(c echo.Context) (*model.Person, error) {

	person := model.Person{}
	if error := c.Bind(&person); error != nil {
		return nil, error
	}
	result := arfaddperson.ARF_AddPerson(&person)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Person{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Person{}) {
		return result.(*model.Person), nil
	} else {
		return nil, nil
	}
}
