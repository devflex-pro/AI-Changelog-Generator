package domain

import "context"

const (
	GitHub = "github"
)

type Commit struct {
	Hash     string
	Author   string
	Commiter string
	Message  string
	//etc...
}
type GitProvider interface {
	FetchCommits(context.Context) ([]Commit, error)
}
