package server

import (
	"reflect"
	"sync"
)

type service struct {
	name    string
	methods map[string]*serviceMethod
}

type serviceMethod struct {
	method reflect.Method
}

type serviceMapper struct {
	mutex    sync.Mutex
	services map[string]*service
}

func registerService() error {
	return nil
}

func (mapper *serviceMapper) getService(name string) (*service, *serviceMethod, error) {
	return nil, nil, nil
}
