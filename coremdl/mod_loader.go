package coremdl

import "strings"

type ModLoader uint32

const (
	// ModLoaderNone is the default value for a mod loader
	ModLoaderNone    ModLoader = 0
	ModLoaderUnknown ModLoader = 1 << iota
	ModLoaderFabric
	ModLoaderForge
	ModLoaderQuilt
	ModLoaderBukkit
	ModLoaderSponge
	ModLoaderSpigot
	ModLoaderPaper
)

func (m ModLoader) Has(other ModLoader) bool {
	return m == ModLoaderNone || m&other != 0
}

func (m ModLoader) String() string {
	switch m {
	case ModLoaderFabric:
		return "fabric"
	case ModLoaderForge:
		return "forge"
	case ModLoaderQuilt:
		return "quilt"
	case ModLoaderBukkit:
		return "bukkit"
	case ModLoaderSponge:
		return "sponge"
	case ModLoaderSpigot:
		return "spigot"
	case ModLoaderPaper:
		return "paper"
	case ModLoaderNone:
		return ""
	default:
		return "unknown"
	}
}

func NewModLoaderFromString(s string) ModLoader {
	s = strings.ToLower(s)

	switch s {
	case "fabric":
		return ModLoaderFabric
	case "forge":
		return ModLoaderForge
	case "quilt":
		return ModLoaderQuilt
	case "bukkit":
		return ModLoaderBukkit
	case "sponge":
		return ModLoaderSponge
	case "spigot":
		return ModLoaderSpigot
	case "paper":
		return ModLoaderPaper
	case "":
		return ModLoaderNone
	default:
		return ModLoaderUnknown
	}
}
