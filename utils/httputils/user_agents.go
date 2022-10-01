package httputils

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

const (
	latestUserAgentsURL = "https://jnrbsn.github.io/user-agents/user-agents.json"
)

// UserAgents used for getting the latest browser user agents
type UserAgents interface {
	GetLatest(ctx context.Context) ([]string, error)
	GetRandom(ctx context.Context) (string, error)
}

type userAgentsImpl struct {
	client *http.Client
	cache  []string
}

func NewLatestUserAgents(client *http.Client) UserAgents {
	return &userAgentsImpl{
		client: client,
	}
}

func (l *userAgentsImpl) GetLatest(ctx context.Context) ([]string, error) {
	if l.cache != nil {
		return l.cache, nil
	}

	l.client = &http.Client{}
	resp, err := l.client.Get(latestUserAgentsURL)
	CloseBody(ctx, resp)
	if err != nil {
		return nil, err
	}

	var userAgents []string
	if err := json.NewDecoder(resp.Body).Decode(&userAgents); err != nil {
		return nil, err
	}

	l.cache = userAgents

	return userAgents, nil
}

func (l *userAgentsImpl) GetRandom(ctx context.Context) (string, error) {
	userAgents, err := l.GetLatest(ctx)
	if err != nil {
		return "", err
	}

	//nolint:gosec // Not a security issue
	return userAgents[rand.Intn(len(userAgents))], nil
}
