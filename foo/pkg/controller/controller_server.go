package controller

import (
	api "foo/pkg/api"
	"github.com/golang/protobuf/ptypes/empty"
)

import (
	"context"
	"foo/pub"
)

type Server struct{}

func (s *Server) AddPerson(ctx context.Context, parameters *pub.AddPersonParameters) (*pub.Person, error) {

	result, err := api.AddPerson(parameters.Person)
	return result, err
}
func (s *Server) GetAllPersons(ctx context.Context, in *empty.Empty) (*pub.ListPersonResponse, error) {

	result, err := api.GetAllPersons()
	var val pub.ListPersonResponse
	val.Person = result
	return &val, err
}
