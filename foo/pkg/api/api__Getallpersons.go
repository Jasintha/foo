package api

import (
	arfgetallpersons "foo/pkg/rule/ARF_GetAllPersons"
	"reflect"
)

import (
	"foo/pkg/model"
	"foo/pub"
	"github.com/jinzhu/copier"
)

func GetAllPersons() ([]*pub.Person, error) {

	result := arfgetallpersons.ARF_GetAllPersons()
	if reflect.TypeOf(result) == reflect.TypeOf([]*model.Person{}) || reflect.TypeOf(result) == reflect.TypeOf(&[]model.Person{}) {
		var data []*pub.Person
		for _, val := range result.([]*model.Person) {
			var d pub.Person
			copier.Copy(&data, &val)
			data = append(data, &d)

		}
		return data, nil
	} else {
		var data []*pub.Person
		return data, nil
	}
}
