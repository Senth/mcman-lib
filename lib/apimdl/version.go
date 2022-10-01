package apimdl

import "github.com/Senth/mcman-lib/lib/coremdl"

type VersionLatestRequest struct {
	ModID     string            `json:"modId"`
	MCVersion coremdl.MCVersion `json:"mcVersion"`
	ModLoader coremdl.ModLoader `json:"modLoader"`
}

type VersionResponse struct {
	LatestVersion coremdl.Version `json:"versions"`
}
