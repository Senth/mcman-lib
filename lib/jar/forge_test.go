package jar

import (
	"fmt"
	"testing"

	"github.com/Senth/mcman-lib/lib/coremdl"
	"github.com/Senth/mcman-lib/lib/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestForgeParser_Parse(t *testing.T) {
	type testData struct {
		t           *testing.T
		input       []byte
		expected    *coremdl.JarInfo
		expectedErr error
	}

	testCases := []struct {
		name      string
		prepareFn func(d *testData)
	}{
		{
			name: "Valid forge mod",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "forge-valid.toml")
				d.expected = &coremdl.JarInfo{
					NameID:        "jei",
					ModLoaders:    coremdl.ModLoaderForge,
					Name:          "Just Enough Items",
					Description:   "JEI is an item and recipe viewing mod for Minecraft, built from the ground up for stability and performance.",
					VersionNumber: "7.6.4.86",
				}
			},
		},
		{
			name: "Success when containing inline comment",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "forge-invalid-multiline-string.toml")
				d.expected = &coremdl.JarInfo{
					NameID:        "twilightforest",
					ModLoaders:    coremdl.ModLoaderForge,
					Name:          "The Twilight Forest",
					Description:   "An enchanted forest dimension.",
					VersionNumber: "${file.jarVersion}",
				}
			},
		},
		{
			name: "Error when missing mods section",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "forge-invalid-missing-mods.toml")
				d.expectedErr = fmt.Errorf("")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Assemble
			d := &testData{t: t}
			tc.prepareFn(d)

			// Act
			actual, actualErr := newForgeParser().Parse(d.input)

			// Assert
			assert.Equal(t, d.expected, actual)
			assert.IsType(t, d.expectedErr, actualErr)
		})
	}
}
