package template

import (
	"bytes"
	"html/template"
	"strings"
)

func RenderString(templateString string, templateValues interface{}) (string, error) {
	t, err := template.New("").Parse(strings.Join([]string{
		"{{define \"T\"}}",
		templateString,
		"{{end}}",
	}, ""))
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = t.ExecuteTemplate(&out, "T", templateValues); err != nil {
		return "", err
	}

	return out.String(), nil
}

func RenderBytes(templateBytes []byte, templateValues interface{}) ([]byte, error) {
	t, err := template.New("").Parse(strings.Join([]string{
		"{{define \"T\"}}",
		string(templateBytes),
		"{{end}}",
	}, ""))
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err = t.ExecuteTemplate(&out, "T", templateValues); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
