// Package github provides a Go API for the GitHub issue tracker.
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number        int
	HTMLURL       string `json:"html_url"`
	Title         string
	State         string
	User          *User
	CreatedAtTime time.Time `json:"created_at"`
	Body          string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
