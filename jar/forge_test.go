package jar

import (
	"fmt"
	coremdl2 "github.com/Senth/mcman-lib/coremdl"
	"github.com/Senth/mcman-lib/utils/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForgeParser_Parse(t *testing.T) {
	type testData struct {
		t           *testing.T
		input       []byte
		expected    *coremdl2.JarInfo
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
				d.expected = &coremdl2.JarInfo{
					NameID:        "jei",
					ModLoaders:    coremdl2.ModLoaderForge,
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
				d.expected = &coremdl2.JarInfo{
					NameID:        "twilightforest",
					ModLoaders:    coremdl2.ModLoaderForge,
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
