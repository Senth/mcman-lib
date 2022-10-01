package jar

import (
	"io"
	"testing"

	"github.com/Senth/mcman-lib/lib/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestManifestReader_Parse(t *testing.T) {
	type testData struct {
		t           *testing.T
		input       io.ReadCloser
		expected    map[string]string
		expectedErr error
	}

	testCases := []struct {
		name      string
		prepareFn func(d *testData)
	}{
		{
			name: "Valid manifest",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixtureReader(t, "MANIFEST-valid.MF")
				d.expected = map[string]string{
					"Manifest-Version":         "1.0",
					"Implementation-Title":     "s",
					"Implementation-Version":   "4.1.1220",
					"Implementation-Vendor":    "TeamTwilight",
					"Implementation-Timestamp": "2022-07-27T00:47:13+0000",
					"Specification-Vendor":     "TeamTwilight",
					"Specification-Version":    "4.1.1220",
					"Specification-Title":      "Twilight Forest",
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Assemble
			d := &testData{t: t}
			tc.prepareFn(d)
			defer func(reader io.ReadCloser) {
				err := reader.Close()
				if err != nil {
					t.Errorf("Failed to close reader: %v", err)
				}
			}(d.input)

			// Act
			manifest, err := manifestReader{}.Parse(d.input)

			// Assert
			assert.Equal(t, d.expected, manifest)
			assert.IsType(t, d.expectedErr, err)
		})
	}
}
