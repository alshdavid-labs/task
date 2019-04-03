package task

import (
	"fmt"
	"reflect"
)

type Task struct {
	params  []interface{}
	results []interface{}
	fn      interface{}
}

func (bf *Task) Params(params ...interface{}) *Task {
	bf.params = params
	return bf
}

func (bf *Task) Results(results ...interface{}) *Task {
	bf.results = results
	return bf
}

func (bf *Task) Do() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	var params []reflect.Value
	for _, p := range bf.params {
		params = append(params, reflect.ValueOf(p))
	}

	results := reflect.ValueOf(bf.fn).Call(params)

	if bf.results == nil {
		return
	}
	for i, r := range results {
		reflect.ValueOf(bf.results[i]).Elem().Set(r)
	}

	return
}

func New(fn interface{}) *Task {
	return &Task{fn: fn}
}
