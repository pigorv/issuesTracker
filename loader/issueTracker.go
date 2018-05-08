package loader

import (
	"io"
	"github.com/pigorv/issuesTracker/github"
	"log"
	"html/template"
)

var issueList = template.Must(
	template.New("issueList").Parse(
		`<h1>{{.TotalCount}} issues </h1>
		<table>
			<tr style='text-align: left'>
				<th>#</th>
				<th>State</th>
				<th>User</th>
				<th>Title</th>
			</tr>
			{{range .Items}}
			<tr>
				<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
				<td>{{.State}}</td>
				<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
				<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
			</tr>
			{{end}}
		</table>
		`))

// WriteToHtmlPage returns generated html page with search results
func WriteToHtmlPage(searchParams []string, w io.Writer) {
	result, err := github.SearchIssues(searchParams)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	} 
}