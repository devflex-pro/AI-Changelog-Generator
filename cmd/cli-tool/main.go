package main

import (
	"flag"
	"log"

	clitool "github.com/devflex-pro/AI-Changelog-Generator/cli-tool"
	"github.com/devflex-pro/AI-Changelog-Generator/domain"
	git_provider_github "github.com/devflex-pro/AI-Changelog-Generator/git-providers/github"
)

func main() {

	gitProviderName := *flag.String(
		"provider",
		"github",
		"The name of the provider (e.g., github, gitlab)")
	accessToken := *flag.String(
		"token",
		"",
		"Personal access token for the selected provider (required)")
	repoURL := *flag.String(
		"repo",
		"",
		"Repository URL for the selected provider (required)")

	flag.Parse()

	if accessToken == "" || repoURL == "" {
		log.Fatal("Both -token and -repo flags are required")
	}

	var (
		gitProvider domain.GitProvider
		err         error
	)

	switch gitProviderName {
	case domain.GitHub:
		gitProvider, err = git_provider_github.New(accessToken, repoURL)
	default:
		log.Fatalf("Unsupported git provider: %s", gitProviderName)
	}
	if err != nil {
		log.Fatalf("Set git provider failed: %s", err.Error())
	}

	err = clitool.Run(gitProvider)
	if err != nil {
		log.Fatalf("Run failed: %s", err.Error())
	}

}
