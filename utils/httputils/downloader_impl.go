package httputils

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const Timeout = 30 * time.Second

type downloaderImpl struct {
	client    *http.Client
	userAgent string
}

func NewDownloader(userAgent string) Downloader {
	c := &http.Client{
		Timeout: Timeout,
	}
	return &downloaderImpl{
		client:    c,
		userAgent: userAgent,
	}
}

func (d downloaderImpl) Download(ctx context.Context, url string) (b []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	if d.userAgent != "" {
		req.Header.Set("User-Agent", d.userAgent)
	}

	//nolint:bodyclose
	resp, err := d.client.Do(req)
	CloseBody(ctx, resp)
	if err != nil {
		return
	}

	b, err = ReadBody(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	return
}
