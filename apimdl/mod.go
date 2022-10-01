package apimdl

import (
	"fmt"
	coremdl2 "github.com/Senth/mcman-lib/coremdl"
)

type ModFindRequest struct {
	NameID     string             `json:"nameId"`
	Slugs      []string           `json:"slugs"`
	ModLoaders coremdl2.ModLoader `json:"modLoaders"`
}

type ModSearchRequest struct {
	Query      string             `json:"query"`
	ModLoaders coremdl2.ModLoader `json:"modLoaders"`
	MCVersion  coremdl2.MCVersion `json:"mcVersion"`
}

type ModSearchResult struct {
	Results []coremdl2.Mod `json:"results"`
}

func (r ModFindRequest) String() string {
	return fmt.Sprintf("%s-%v-%d", r.NameID, r.Slugs, r.ModLoaders)
}

func (r ModSearchRequest) String() string {
	return fmt.Sprintf("%s-%d-%s", r.Query, r.ModLoaders, r.MCVersion)
}
