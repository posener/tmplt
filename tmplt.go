// Package tmplt is a small wrapper around Go templates for handling simple templates.
package tmplt

import (
	html "html/template"
	"io"
	"strings"
	text "text/template"
)

type Template struct {
	exe interface {
		Execute(io.Writer, interface{}) error
	}
}

// Text creates a new text/template template. It may panic so it is better to be used in global
// scope.
func Text(s string) Template {
	return Template{exe: text.Must(text.New("new").Parse(s))}
}

// Text creates a new html/template template. It may panic so it is better to be used in global
// scope.
func HTML(s string) Template {
	return Template{exe: html.Must(html.New("new").Parse(s))}
}

// Execute the template on data and return the result and an error.
func (t Template) Execute(data interface{}) (string, error) {
	var s strings.Builder
	err := t.exe.Execute(&s, data)
	return s.String(), err
}
