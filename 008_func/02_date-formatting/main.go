package main

import (
	"text/template"
	"time"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func dayMonthYear(t time.Time) string {
	return t.Format("02/01/2006 03:04:05PM")
}

var fm = template.FuncMap{
	"fdateDMY": dayMonthYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln()
	}
}
