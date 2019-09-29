package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func PrintEmails() {

	type Person struct {
		Name   string
		Emails []string
	}

	const templ = `The name is {{.Name}}.
	{{range .Emails}}
		An email is "{{. | emailExpand}}"
	{{end}}
	`

	person := Person{
		"Patrice",
		[]string{"patricepetit@hotmail.com", "petitpatrice@gmail.com"},
	}

	t := template.New("Person template")

	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})

	t, err := t.Parse(templ)

	checkError(err)

	err = t.Execute(os.Stdout, person)

	checkError(err)

}

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return substrs[0] + " at " + substrs[1]
}
