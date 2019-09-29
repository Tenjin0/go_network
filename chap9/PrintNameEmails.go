package main

import (
	"os"
	"text/template"
)

func PrintNameEmails() {

	type Person struct {
		Name   string
		Emails []string
	}

	person := Person{
		"patrice",
		[]string{"patricepetit@hotmail.com", "petitpatrice@gmail.com"},
	}

	templ := `{{- $name := .Name -}}
    {{- range .Emails}}Name is {{$name}} with email {{.}}
{{end -}}
`

	t := template.New("Person template")
	t, err := t.Parse(templ)

	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}
