package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
	"github.com/pigorv/issuesTracker/loader"
)

const mainPage = `
	<div>
		<form action='/search' method='get'>
			Search Issues:<br>
  			<input type="text" name="q"><br>
  			<input type="submit" value="Submit">
		</form>
	</div>`

func main()  {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, mainPage)
}

func searchHandler(w http.ResponseWriter, r *http.Request)  {
	loader.WriteToHtmlPage(strings.Split(r.URL.RawQuery, "%20"), w)
}