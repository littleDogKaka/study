package main

import (
	"fmt"
	"reflect"
	"study/src/tlib"
)

var s = tlib.New("12")

func main() {

	set := reflect.ValueOf(s).MethodByName("Set")
	v := []reflect.Value{reflect.ValueOf("10")}
	set.Call(v)

	run := reflect.ValueOf(s).MethodByName("Run")
	run.Call([]reflect.Value{})

	get := reflect.ValueOf(s).MethodByName("Get")
	ret := get.Call([]reflect.Value{})
	i := ret[0].Int()
	fmt.Printf("%+v int:%d\n", ret, i)
}
