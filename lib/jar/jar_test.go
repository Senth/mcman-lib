package jar

import (
	"context"
	"fmt"
	"testing"

	"github.com/Senth/mcman-lib/lib/coremdl"
	"github.com/Senth/mcman-lib/lib/utils/test"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestJarImpl_GetMod(t *testing.T) {
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
			name: "Valid fabric mod",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "fabric-valid.jar")
				d.expected = &coremdl.JarInfo{
					NameID:        "carpet",
					ModLoaders:    coremdl.ModLoaderFabric,
					Name:          "Carpet Mod in Fabric",
					Description:   "Carpet made out of fabric",
					VersionNumber: "1.4.16",
				}
			},
		},
		{
			name: "Valid forge mod",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "forge-valid.jar")
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
			name: "${file.jarVersion} to version number",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "file-version-valid.jar")
				d.expected = &coremdl.JarInfo{
					NameID:        "twilightforest",
					ModLoaders:    coremdl.ModLoaderForge,
					Name:          "The Twilight Forest",
					Description:   "An enchanted forest dimension.",
					VersionNumber: "4.1.1220",
				}
			},
		},
		{
			name: "${file.jarVersion} success when missing specification version",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "file-version-missing-specification.jar")
				d.expected = &coremdl.JarInfo{
					NameID:        "twilightforest",
					ModLoaders:    coremdl.ModLoaderForge,
					Name:          "The Twilight Forest",
					Description:   "An enchanted forest dimension.",
					VersionNumber: "4.1.1220",
				}
			},
		},
		{
			name: "${file.jarVersion} success when missing implementation version",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "file-version-missing-implementation.jar")
				d.expected = &coremdl.JarInfo{
					NameID:        "twilightforest",
					ModLoaders:    coremdl.ModLoaderForge,
					Name:          "The Twilight Forest",
					Description:   "An enchanted forest dimension.",
					VersionNumber: "4.1.1220",
				}
			},
		},
		{
			name: "${file.jarVersion} error when missing both specification and implementation version",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "file-version-missing-version.jar")
				d.expectedErr = fmt.Errorf("")
			},
		},
		{
			name: "${file.jarVersion} error when missing MANIFEST.MF file",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "file-version-missing-MANIFEST.jar")
				d.expectedErr = fmt.Errorf("")
			},
		},
		{
			name: "Invalid jar file",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "invalid.jar")
				d.expected = nil
				d.expectedErr = fmt.Errorf("")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Assemble
			d := &testData{
				t: t,
			}
			tc.prepareFn(d)

			// Act
			actual, actualErr := NewJar().GetMod(ctx, d.input)

			// Assert
			assert.Equal(t, d.expected, actual)
			assert.IsType(t, d.expectedErr, actualErr)
		})
	}
}
