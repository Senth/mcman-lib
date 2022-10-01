package jar

import (
	"bufio"
	"io"
	"strings"
)

const maxSplits = 2

type manifestReader struct{}

func (m manifestReader) Parse(reader io.Reader) (map[string]string, error) {
	s := bufio.NewScanner(reader)

	manifest := make(map[string]string)

	for s.Scan() {
		line := s.Text()
		fields := strings.SplitN(line, ": ", maxSplits)

		if len(fields) != maxSplits {
			continue
		}

		manifest[fields[0]] = fields[1]
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return manifest, nil
}
