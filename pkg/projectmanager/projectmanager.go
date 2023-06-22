package projectmanager

import (
	"path/filepath"

	"github.com/alomia/fastman/pkg/fileutils"
	"github.com/alomia/fastman/pkg/sampledata"
)

type Package struct {
	Name  string
	Files []string
}

type ProjectStructure struct {
	Path     string
	Packages []Package
	Files    []string
}

func NewProjectStructure(path string, packages []Package, files []string) *ProjectStructure {
	return &ProjectStructure{
		Path:     path,
		Packages: packages,
		Files:    files,
	}
}

func (p *ProjectStructure) CreateProjectStructure() error {
	err := fileutils.CreateConfigFile(p.Path)
	if err != nil {
		return err
	}

	for _, pkg := range p.Packages {
		pkgPath := filepath.Join(p.Path, pkg.Name)

		err := fileutils.CreatePackage(pkgPath)
		if err != nil {
			return err
		}

		for _, file := range pkg.Files {
			filePath := filepath.Join(pkgPath, file)
			fileContentPath := filepath.Join(pkg.Name, file)

			content, err := sampledata.GetSampleContent(fileContentPath)
			if err != nil {
				return err
			}

			err = fileutils.CreateFile(filePath, []byte(content))
			if err != nil {
				return err
			}
		}
	}

	for _, file := range p.Files {
		filePath := filepath.Join(p.Path, file)

		content, err := sampledata.GetSampleContent(file)
		if err != nil {
			return err
		}

		err = fileutils.CreateFile(filePath, []byte(content))
		if err != nil {
			return err
		}
	}

	return nil
}
