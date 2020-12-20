package controller

import (
	"bytes"
	"testing"
	"text/template"
)

func TestTemplates(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}
	// sweaters := Inventory{"wool", 17}
	list := [4]int{1, 2, 3, 4}
	tmpl, err := template.New("test").Parse("{{range .}}{{.}}{{end}}")
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, list)
	if err != nil {
		panic(err)
	}

	t.Fatal(buf)
}
