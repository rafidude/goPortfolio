package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//go:generate go run gen.go
//go:generate go fmt codegen/user.gen.go

func main() {
	fmt.Println("Hello, world!")

	var resp map[string]interface{}
	in, err := os.Open("./user.json")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &resp)
	fmt.Println(string(b))
	fmt.Println(resp)

	data := struct {
		Name   string
		Fields map[string]interface{}
	}{
		"User",
		resp,
	}

	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"Title": strings.Title,
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")
	fmt.Println(tpl)
	if err != nil {
		panic(err)
	}

	out, _ := os.Create("./out/user.gen.go")
	defer out.Close()

	tpl.Execute(out, data)
}
