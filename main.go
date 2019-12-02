package main

import (
	"os"
	"reflect"
)

type Days struct{}

func main() {
	days := Days{}
	day := reflect.ValueOf(days).MethodByName(os.Args[1])
	day.Call(nil)
}