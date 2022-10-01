package httputils

import "context"

type Downloader interface {
	Download(ctx context.Context, url string) ([]byte, error)
}
