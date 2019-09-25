package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}.
The age is {{.Age -}}.
{{- range .Emails }}
    An Email is {{. -}}
{{end}}
The Jobs are:
{{with .Jobs -}}
{{range .}}
    An employer is {{.Employer}}
    and the role is {{.Role}}
{{end -}}
{{end -}}
`

func PrintPerson() {

	job1 := Job{
		Employer: "Box Hill institute",
		Role:     "Director, commerce and ICT",
	}

	job2 := Job{
		Employer: "Canberra University",
		Role:     "Adjunct Professor",
	}

	person := Person{
		Name:   "Patrice",
		Age:    30,
		Emails: []string{"petitpatrice@gmail.com", "patricepetit@hotmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")

	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)

}
