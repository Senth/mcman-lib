package apimdl

import (
	coremdl2 "github.com/Senth/mcman-lib/coremdl"
)

type VersionLatestRequest struct {
	ModID     string             `json:"modId"`
	MCVersion coremdl2.MCVersion `json:"mcVersion"`
	ModLoader coremdl2.ModLoader `json:"modLoader"`
}

type VersionLatestResponse struct {
	LatestVersion coremdl2.Version `json:"versions"`
}
