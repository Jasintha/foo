package controller_echo

import (
	apiecho "foo/pkg/api_echo"
	model "foo/pkg/model"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func AddPerson(c echo.Context) error {

	result, _ := apiecho.AddPerson(c)
	return c.JSON(http.StatusOK, result)
}

type ListPersonResponse struct {
	Person []*model.Person
}

func GetAllPersons(c echo.Context) error {

	result, _ := apiecho.GetAllPersons(c)
	value := ListPersonResponse{}
	value.Person = result
	return c.JSON(http.StatusOK, value)
}
func APIControllerEcho(g *echo.Group) {

	g.POST("/api/person", AddPerson)
	g.GET("/api/people", GetAllPersons)
}
