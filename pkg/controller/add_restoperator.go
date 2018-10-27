package controller

import (
	"github.com/example-inc/rest-operator/pkg/controller/restoperator"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, restoperator.Add)
}
