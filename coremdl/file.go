package coremdl

type File struct {
	Filename    string `json:"filename"`
	DownloadURL string `json:"downloadUrl"`
	Size        int    `json:"size"`
}
