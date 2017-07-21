package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
