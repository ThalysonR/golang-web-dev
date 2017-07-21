package main

import (
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip String
}

type region struct {
	Term string
	Hotels []hotel
}

type geral struct {
	Southern region
	Central region
	Northern region
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

}
