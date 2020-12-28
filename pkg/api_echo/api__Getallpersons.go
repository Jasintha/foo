package api_echo

import (
	model "foo/pkg/model"
	arfgetallpersons "foo/pkg/rule/ARF_GetAllPersons"
	"reflect"
)

import (
	"github.com/labstack/echo/v4"
)

func GetAllPersons(c echo.Context) ([]*model.Person, error) {

	result := arfgetallpersons.ARF_GetAllPersons()
	if reflect.TypeOf(result) == reflect.TypeOf([]*model.Person{}) || reflect.TypeOf(result) == reflect.TypeOf([]*model.Person{}) {
		return result.([]*model.Person), nil
	} else {
		return nil, nil
	}
}
