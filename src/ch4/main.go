package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	report := template.Must(template.New("Issue list").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

	if err := report.Execute(os.Stdout, &result); err != nil {
		log.Fatal(err)
	}

}
