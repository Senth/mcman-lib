package coremdl

import "time"

type Stability string

const (
	StabilityAlpha   Stability = "alpha"
	StabilityBeta    Stability = "beta"
	StabilityRelease Stability = "release"
	StabilityUnknown Stability = "unknown"
)

type Version struct {
	SiteVersionID string       `json:"siteVersionId" firestore:"siteVersionId"`
	ModID         string       `json:"modId" firestore:"modId"`
	SiteModID     string       `json:"siteModId" firestore:"siteModId"`
	Number        string       `json:"number" firestore:"number"`
	Stability     Stability    `json:"stability" firestore:"stability"`
	SiteType      SiteType     `json:"siteType" firestore:"siteType"`
	UploadTime    time.Time    `json:"uploadTime" firestore:"uploadTime"`
	MCVersions    MCVersions   `json:"minecraftVersions" firestore:"minecraftVersions"`
	File          File         `json:"file" firestore:"file"`
	ModLoader     ModLoader    `json:"modLoader" firestore:"modLoader"`
	Dependencies  []Dependency `json:"dependencies" firestore:"dependencies"`
	Timestamp     time.Time    `json:"timestamp" firestore:"timestamp"`
}

type MCVersion string

func (v MCVersion) String() string {
	return string(v)
}

func (v MCVersion) IsSet() bool {
	return v != ""
}

type MCVersions []MCVersion

func (v MCVersions) Contains(version MCVersion) bool {
	for _, v := range v {
		if v == version {
			return true
		}
	}

	return false
}
