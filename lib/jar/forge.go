package jar

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Senth/mcman-lib/lib/coremdl"
	"github.com/Senth/mcman-lib/lib/utils/ptr"
)

type forgeParser struct{}

func newForgeParser() parser {
	return &forgeParser{}
}

func (f forgeParser) Parse(data []byte) (*coremdl.JarInfo, error) {
	data = f.cleanData(data)

	t := forgeTOML{}

	err := toml.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	if t.Mods == nil || len(t.Mods) == 0 {
		return nil, fmt.Errorf("no mods found in forge mod file")
	}

	return ptr.New(t.Model()), nil
}

func (f forgeParser) cleanData(data []byte) []byte {
	data = f.fixUnescapedNewlines(data)
	return data
}

var unescapedNewlinesTOMLRegexp = regexp.MustCompile(`= ?"([^"]*\n)`)

func (f forgeParser) fixUnescapedNewlines(data []byte) []byte {
	matches := unescapedNewlinesTOMLRegexp.FindAllSubmatch(data, -1)
	for _, match := range matches {
		replaced := bytes.ReplaceAll(match[1], []byte("\n"), []byte("\\n"))
		data = bytes.Replace(data, match[1], replaced, 1)
	}

	return data
}

type forgeTOML struct {
	Mods []struct {
		ModID       string `toml:"modId"`
		Version     string `toml:"version"`
		DisplayName string `toml:"displayName"`
		Description string `toml:"description"`
	} `toml:"mods"`
}

func (f forgeTOML) Model() coremdl.JarInfo {
	return coremdl.JarInfo{
		NameID:        f.Mods[0].ModID,
		ModLoaders:    coremdl.ModLoaderForge,
		Name:          f.Mods[0].DisplayName,
		Description:   strings.Trim(f.Mods[0].Description, " \t\n"),
		VersionNumber: f.Mods[0].Version,
	}
}
