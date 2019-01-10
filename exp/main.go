package main

import (
	"html/template"
	"os"
)

// User type for the data
type User struct {
	Name   string
	Dog    string
	Lists  []string
	Arrays map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "John Smith",
		Dog:   "smiluy",
		Lists: []string{"a", "b", "c"},
		Arrays: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
