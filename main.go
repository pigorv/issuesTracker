package main

import (
	"strings"
	"time"
	"fmt"
	"github.com/pigorv/issuesTracker/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	var currentMonthCreated []string
	timeNow := time.Now()
	for _, item := range result.Items {
		if item.CreatedAt.Month() == timeNow.Month() && item.CreatedAt.Year() == timeNow.Year() {
			line := fmt.Sprintf("#%-5d %9.9s %.55s Created:%s", item.Number, item.User.Login, item.Title, item.CreatedAt)
			currentMonthCreated = append(currentMonthCreated, line)
		}
		println(strings.Join(currentMonthCreated,"\n"))
	}

}