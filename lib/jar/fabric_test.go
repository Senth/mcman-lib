package jar

import (
	"testing"

	"github.com/Senth/mcman-lib/lib/coremdl"
	"github.com/Senth/mcman-lib/lib/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestFabricParser_Parse(t *testing.T) {
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
				d.input = test.LoadFixture(t, "fabric-valid.json")
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
			name: "Success when containing invalid character",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "fabric-invalid-character.json")
				d.expected = &coremdl.JarInfo{
					NameID:        "capes",
					ModLoaders:    coremdl.ModLoaderFabric,
					Name:          "Capes",
					Description:   "Check needs his Cape",
					VersionNumber: "1.1.2",
				}
			},
		},
		{
			name: "Success when containing invalid control character",
			prepareFn: func(d *testData) {
				d.input = test.LoadFixture(t, "fabric-invalid-control-character.json")
				d.expected = &coremdl.JarInfo{
					NameID:        "itemmodelfix",
					ModLoaders:    coremdl.ModLoaderFabric,
					Name:          "Item\nModel\nFix",
					Description:   "Fixes gaps in generated item models.\nTo access the configuration, follow the instructions on the mod website.",
					VersionNumber: "1.0.2+1.17",
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Assemble
			d := &testData{t: t}
			tc.prepareFn(d)

			// Act
			actual, actualErr := newFabricParser().Parse(d.input)

			// Assert
			assert.Equal(t, d.expected, actual)
			assert.IsType(t, d.expectedErr, actualErr)
		})
	}
}
