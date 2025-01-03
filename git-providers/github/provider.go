package git_provider_github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/devflex-pro/AI-Changelog-Generator/domain"
)

type provider struct {
	token string
	owner string
	repo  string
}

func New(token, repoURL string) (*provider, error) {
	owner, repo, err := parseGitHubURL(repoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GitHub URL: %w", err)
	}

	return &provider{
		token: token,
		owner: owner,
		repo:  repo,
	}, nil
}

func (p *provider) FetchCommits(ctx context.Context) ([]domain.Commit, error) {

	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/commits",
		p.owner,
		p.repo,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+p.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commits: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResponse []GitHubCommitResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var commits []domain.Commit
	for _, c := range apiResponse {
		commits = append(commits, domain.Commit{
			Hash:     c.SHA,
			Author:   c.Commit.Author.Name,
			Commiter: c.Commit.Committer.Name,
			Message:  c.Commit.Message,
		})
	}

	return commits, nil
}

func parseGitHubURL(url string) (owner, repo string, err error) {
	parts := strings.Split(strings.TrimPrefix(url, "https://github.com/"), "/")
	if len(parts) < 2 {
		return "", "", errors.New("invalid GitHub URL format")
	}
	return parts[0], parts[1], nil
}
