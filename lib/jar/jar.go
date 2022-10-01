package jar

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/Senth/mcman-lib/lib/coremdl"
)

type Jar interface {
	GetMod(ctx context.Context, data []byte) (*coremdl.JarInfo, error)
}

type parser interface {
	Parse(data []byte) (*coremdl.JarInfo, error)
}

type jarImpl struct {
	forge    parser
	fabric   parser
	manifest manifestReader
}

func NewJar() Jar {
	return &jarImpl{
		forge:    newForgeParser(),
		fabric:   newFabricParser(),
		manifest: manifestReader{},
	}
}

func (j jarImpl) GetMod(ctx context.Context, data []byte) (*coremdl.JarInfo, error) {
	// Read zip/jar
	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	// Get zip mod file information
	file, modLoader, err := j.getZippedFile(zipReader)
	if err != nil {
		return nil, err
	}

	// Open zip file
	fileReader, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer func(fileReader io.ReadCloser) {
		_ = fileReader.Close()
	}(fileReader)

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(fileReader)
	if err != nil {
		return nil, err
	}

	// Parse json with mod information
	info, err := j.parseJar(ctx, modLoader, buf.Bytes())
	if err != nil {
		return nil, err
	}

	// Fix version number if ${file.jarVersion} is used
	if info.VersionNumber == "${file.jarVersion}" {
		info.VersionNumber, err = j.getJARVersion(zipReader)
		if err != nil {
			return nil, err
		}
	}

	return info, err
}

func (j jarImpl) parseJar(ctx context.Context, modLoader coremdl.ModLoader, data []byte) (*coremdl.JarInfo, error) {
	// Parse json with mod information
	switch modLoader {
	case coremdl.ModLoaderForge:
		return j.forge.Parse(data)
	case coremdl.ModLoaderFabric:
		return j.fabric.Parse(data)
	default:
		return nil, fmt.Errorf("unknown mod loader %s", modLoader)
	}
}

func (j jarImpl) getZippedFile(r *zip.Reader) (*zip.File, coremdl.ModLoader, error) {
	for _, f := range r.File {
		switch f.Name {
		case "fabric.mod.json":
			return f, coremdl.ModLoaderFabric, nil
		case "META-INF/mods.toml":
			return f, coremdl.ModLoaderForge, nil
		}
	}

	return nil, coremdl.ModLoaderNone, fmt.Errorf("could not find mod file in jar")
}

func (j jarImpl) getJARVersion(r *zip.Reader) (string, error) {
	var file *zip.File

	for _, f := range r.File {
		if f.Name == "META-INF/MANIFEST.MF" {
			file = f
			break
		}
	}

	if file == nil {
		return "", fmt.Errorf("could not find manifest file in jar")
	}

	// Open zip file
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}

	defer func(fileReader io.ReadCloser) {
		err = fileReader.Close()
	}(fileReader)

	// Parse manifest file
	manifest, err := j.manifest.Parse(fileReader)
	if err != nil {
		return "", err
	}

	// Use implementation version if available
	version := manifest["Implementation-Version"]
	if version != "" {
		return version, err
	}

	// Use specification version if available
	version = manifest["Specification-Version"]
	if version == "" {
		return "", fmt.Errorf("could not find version in manifest")
	}

	return version, err
}
