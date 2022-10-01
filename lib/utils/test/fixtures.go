package test

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// LoadFixtureJSON from testdata directory and unmarshal the JSON into the given interface
func LoadFixtureJSON(t *testing.T, filename string, data any) {
	t.Helper()

	b := LoadFixture(t, filename)

	err := json.Unmarshal(b, data)
	if err != nil {
		t.Fatal(err)
	}
}

// LoadFixture from testdata directory as []byte
func LoadFixture(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatal(err)
	}

	return b
}

// LoadFixtureReader from testdata directory as io.ReadCloser
func LoadFixtureReader(t *testing.T, filename string) io.ReadCloser {
	t.Helper()

	f, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatal(err)
	}

	return f
}
