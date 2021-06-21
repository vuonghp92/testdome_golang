package main

import (
	"fmt"
	"reflect"
)

func uniqueNames(a, b []string) []string {
	var result []string
	for _, val := range a {
		exists, _ := in_array(val, result)
		if !exists {
			result = append(result, val)
		}
	}

	for _, val := range b {
		exists, _ := in_array(val, result)
		if !exists {
			result = append(result, val)
		}
	}

	return result
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
