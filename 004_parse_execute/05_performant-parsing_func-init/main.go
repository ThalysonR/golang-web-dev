package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	check(err)
}