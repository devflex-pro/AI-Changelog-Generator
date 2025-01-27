package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	clitool "github.com/devflex-pro/AI-Changelog-Generator/cli-tool"
	"github.com/devflex-pro/AI-Changelog-Generator/domain"
	git_provider_github "github.com/devflex-pro/AI-Changelog-Generator/git-providers/github"
)

func main() {

	accessToken := flag.String(
		"token",
		"",
		"Personal access token for the selected provider (required)")
	repoURL := flag.String(
		"repo",
		"",
		"Repository URL for the selected provider (required)")

	flag.Parse()
	fmt.Printf("Repo: %s, Token: %s\n", *repoURL, *accessToken)
	if *accessToken == "" || *repoURL == "" {
		log.Fatal("Both -token and -repo flags are required")
	}

	var (
		gitProvider domain.GitProvider
		err         error
	)

	parsedRepoURL, err := url.Parse(*repoURL)
	if err != nil {
		log.Fatalf("Parse repo URl failed: %s", err.Error())
	}

	switch parsedRepoURL.Hostname() {
	case domain.GitHub:
		gitProvider, err = git_provider_github.New(*accessToken, *repoURL)
		if err != nil {
			log.Fatalf("GitHub provider init failed: %s", err.Error())
		}
	default:
		log.Fatalf("Unsupported git provider: %s", parsedRepoURL.Hostname())
	}
	if err != nil {
		log.Fatalf("Set git provider failed: %s", err.Error())
	}

	err = clitool.Run(gitProvider)
	if err != nil {
		log.Fatalf("Run failed: %s", err.Error())
	}

}
