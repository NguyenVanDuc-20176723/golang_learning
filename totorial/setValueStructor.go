package main

import (
	"fmt"
	"reflect"
)

type Paging struct {
	NextPage string `q:"next_page"`
	PrevPage string `q:"prev_page"`
}

type PagingMeta struct {
	NextPage string `q:"next_page2"`
	PrevPage string `q:"prev_page2"`
}

type Response struct {
	Paging
	Name string `q:"name"`
	Age  int    `q:"age"`
	Meta PagingMeta
}

type Request struct {
	NextPage string `q:"next_page"`
	PrevPage string `q:"prev_page"`
	Name     string `q:"name"`
}

func reUpdate(req interface{}, params map[string]interface{}) {
	v := reflect.ValueOf(req).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct {
			reUpdate(v.Field(i).Interface(), params)
		}
		field := t.Field(i)
		tagName := field.Tag.Get("q")
		if tagName == "" {
			continue
		}
		fieldValue := reflect.ValueOf(params[tagName])
		if fieldValue.Type().AssignableTo(v.Field(i).Type()) {
			v.Field(i).Set(fieldValue)
		}
	}
}

func updateRequest(req interface{}, params map[string]interface{}) {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	} else if v.Kind() == reflect.Struct {
		v = reflect.ValueOf(req)
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct {
			value := v.Field(i)
			for j := 0; j < value.NumField(); j++ {
				field := value.Type().Field(j)
				tagName := field.Tag.Get("q")
				if tagName == "" {
					continue
				}
				val, ok := params[tagName]
				if !ok {
					continue
				}
				fieldValue := reflect.ValueOf(val)
				if fieldValue.Type().AssignableTo(value.Field(j).Type()) {
					value.Field(j).Set(fieldValue)
				}
			}
		}

		field := t.Field(i)
		tagName := field.Tag.Get("q")
		if tagName == "" {
			continue
		}
		fieldValue := reflect.ValueOf(params[tagName])
		if fieldValue.Type().AssignableTo(v.Field(i).Type()) {
			v.Field(i).Set(fieldValue)
		}
	}
	return
}

func main() {
	data := map[string]interface{}{
		"next_page":  "123",
		"prev_page":  "234",
		"next_page2": "1234",
		"prev_page2": "2345",

		"name": "duc",
		"age":  23,
	}

	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//
	//}
	var resp Response
	//err = json.Unmarshal(jsonData, &resp)

	//jsonData, err = json.Marshal(resp)
	//if err != nil {
	//
	//}
	updateRequest(&resp, data)
	fmt.Println()
}
