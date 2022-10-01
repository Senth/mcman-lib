package coremdl

type JarInfo struct {
	NameID        string    `json:"nameId"`
	ModLoaders    ModLoader `json:"modLoaders"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	VersionNumber string    `json:"version"`
}
