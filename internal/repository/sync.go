package repository

import (
	"context"
	"net/http"

	"github.com/google/go-github/v71/github"
	"golang.org/x/oauth2"
)

// Syncer is the interface for syncing information from a repository
type Syncer interface {
	Sync(ctx context.Context, owner, repo string) error
}

type syncer struct {
	client github.Client
}

func New(ctx context.Context, token string) Syncer {
	return &syncer{
		client: *newGithubClient(ctx, token),
	}
}

func newGithubClient(ctx context.Context, token string) *github.Client {
	var httpClient *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(ctx, ts)
	}

	return github.NewClient(httpClient)
}

func (s *syncer) Sync(ctx context.Context, owner, repo string) error {
	return nil
}
