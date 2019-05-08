package test
import (
	"testing"
	"reflect"
)

func TestResultFunc(t *testing.T) {

	call1 := func(v1 int,v2 int,v3 string){
		t.Log(v1,v2,v3)
	}

	call2 := func(v1 int,v2 int){
		t.Log(v1,v2)
	}

	var (
		function reflect.Value
		value []reflect.Value
		n int
	)

	bridge := func(call interface{},args... interface{}){
		n = len(args)
		value = make([]reflect.Value,n)
		for i := 0;i < n;i ++{
			value[i] = reflect.ValueOf(args[i])
		}

		function = reflect.ValueOf(call)
		function.Call(value)
	}
	bridge(call2,1,2)
	bridge(call1,1,2,"12")
}