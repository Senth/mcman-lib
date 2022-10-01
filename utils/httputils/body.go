package httputils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// ReadBody reads the body of a response and returns it as a byte slice.
func ReadBody(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}

// ReadBodyJSON reads the body of a response and unmarshals the JSON into the given interface
func ReadBodyJSON(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}

func CloseBody(ctx context.Context, resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
}
