package coremdl

type Dependency struct {
	VersionID string         `json:"versionId"`
	SiteModID string         `json:"siteModId"`
	Type      DependencyType `json:"type"`
	SiteType  SiteType       `json:"siteType"`
	Mod       *Mod           `json:"mod,omitempty"`
}

type DependencyType string

const (
	DependencyTypeRequired     DependencyType = "required"
	DependencyTypeOptional     DependencyType = "optional"
	DependencyTypeIncompatible DependencyType = "incompatible"
	DependencyTypeUnknown      DependencyType = "unknown"
)
