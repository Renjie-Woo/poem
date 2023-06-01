package main

import (
	"doraemon/poem/demo"
	"fmt"
	"reflect"
)

//func main() {
//	fmt.Println(dao.GetKnowledge())
//}

func Marshal(value interface{}) {
	result := "{"
	elem := reflect.TypeOf(value).Elem()
	elemValue := reflect.ValueOf(value).Elem()
	for i := 0; i < elem.NumField(); i++ {
		tag := elem.Field(i).Tag.Get("my_json")
		if tag == "" {
			continue
		}
		result += "\"" + tag + "\":"
		v := elemValue.Field(i).Elem()
		result += fmt.Sprintf("%v,", v)
	}
	result = result[:len(result)-1] + "}"
	fmt.Println(result)
}

func main() {
	e := &demo.Event{}
	e.SetValue("hello", 12)
	Marshal(e)
}
