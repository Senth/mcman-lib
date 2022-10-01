package coremdl

import (
	"time"
)

// Mod contains internal ID including information about the mod on different sites.
type Mod struct {
	ID        string    `json:"id" firestore:"id"`
	NameID    string    `json:"nameId" firestore:"nameId"`
	Sites     ModSites  `json:"sites" firestore:"sites"`
	Timestamp time.Time `json:"timestamp" firestore:"timestamp"`
}

// ModSite mod information for a specific site
type ModSite struct {
	ID          string     `json:"id" firestore:"id"`
	Name        string     `json:"name" firestore:"name"`
	Type        SiteType   `json:"type" firestore:"type"`
	Slug        string     `json:"slug" firestore:"slug"`
	Description string     `json:"description" firestore:"description"`
	ModLoaders  ModLoader  `json:"modLoaders" firestore:"modLoaders"`
	MCVersions  MCVersions `json:"mcVersions" firestore:"mcVersions"`
	Updated     time.Time  `json:"updated" firestore:"updated"`
	Published   time.Time  `json:"published" firestore:"published"`
	Timestamp   time.Time  `json:"timestamp" firestore:"timestamp"` // Timestamp when this information was updated in DB
}

// HasModLoader returns true if the mod has the specified mod loader
// Will also return true if the mod loader is none in the mod site or parameter
func (m ModSite) HasModLoader(modLoader ModLoader) bool {
	if m.ModLoaders == ModLoaderNone || modLoader == ModLoaderNone {
		return true
	}

	return m.ModLoaders&modLoader > 0
}

// SiteType Where the mod can be found
type SiteType string

const (
	SiteTypeCurse    SiteType = "curse"
	SiteTypeModrinth SiteType = "modrinth"
)

// ModSites map of mod sites
type ModSites []ModSite

func (m ModSites) GetType(Type SiteType) (int, *ModSite) {
	for i, modSite := range m {
		if modSite.Type == Type {
			return i, &modSite
		}
	}

	return -1, nil
}
