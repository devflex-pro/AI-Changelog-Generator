package clitool

import (
	"context"
	"fmt"

	"github.com/devflex-pro/AI-Changelog-Generator/domain"
)

func Run(
	gitProvider domain.GitProvider,
) error {

	commits, err := gitProvider.FetchCommits(context.TODO())
	if err != nil {
		return fmt.Errorf(
			"Error fetching commits: %s",
			err.Error(),
		)
	}

	for _, commit := range commits {
		fmt.Printf(
			"Hash: %s, Author: %s, Message: %s\n",
			commit.Hash,
			commit.Author,
			commit.Message,
		)
	}

	return nil
}
