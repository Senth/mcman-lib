package jar

import (
	"bytes"
	"encoding/json"
	coremdl2 "github.com/Senth/mcman-lib/coremdl"
	"github.com/Senth/mcman-lib/utils/ptr"
	"regexp"
)

type fabricParser struct{}

func newFabricParser() parser {
	return &fabricParser{}
}

func (f fabricParser) Parse(data []byte) (*coremdl2.JarInfo, error) {
	data = f.cleanData(data)

	j := fabricJSON{}

	err := json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	return ptr.New(j.Model()), nil
}

func (f fabricParser) cleanData(data []byte) []byte {
	data = f.fixUnescapedNewlines(data)
	return data
}

var unescapedNewlinesJSONRegexp = regexp.MustCompile(`: ?"([^"]*\n)`)

func (f fabricParser) fixUnescapedNewlines(data []byte) []byte {
	matches := unescapedNewlinesJSONRegexp.FindAllSubmatch(data, -1)
	for _, match := range matches {
		replaced := bytes.ReplaceAll(match[1], []byte("\n"), []byte("\\n"))
		data = bytes.Replace(data, match[1], replaced, 1)
	}

	return data
}

type fabricJSON struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (f fabricJSON) Model() coremdl2.JarInfo {
	return coremdl2.JarInfo{
		NameID:        f.ID,
		ModLoaders:    coremdl2.ModLoaderFabric,
		Name:          f.Name,
		Description:   f.Description,
		VersionNumber: f.Version,
	}
}
