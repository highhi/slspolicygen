package slspolicygen

import (
	"os"
	"strings"
	"text/template"
)

func Gen(file *os.File, params map[string]interface{}) error {
	tpl, err := template.New("").Parse(strings.TrimSpace(policyTemplate))

	if err != nil {
		return err
	}

	if err = tpl.Execute(file, params); err != nil {
		return err
	}

	return nil
}
