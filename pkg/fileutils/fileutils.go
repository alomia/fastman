package fileutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/fastman/pkg/sampledata"
)

func CreateFile(path string, content ...[]byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = fmt.Errorf("error closing file: %w", closeErr)
		}
	}()

	if len(content) > 0 && content[0] != nil {
		_, err = file.Write(content[0])
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

func CreateDirectory(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		err = os.MkdirAll(absPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	return nil
}

func CreatePackage(path string) error {
	namePackage, _ := filepath.Abs(path)
	if err := CreateDirectory(path); err != nil {
		return fmt.Errorf("error creating \"%s\" package: %w", namePackage, err)
	}

	initFilePath := filepath.Join(path, "__init__.py")
	if err := CreateFile(initFilePath); err != nil {
		return fmt.Errorf("error creating __init__.py file in package \"%s\": %w", namePackage, err)
	}

	return nil
}

func FileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("error verifying the existence of the file or directory: %v", err)
}

func CreateConfigFile(currentDir string) error {
	configFileName := "fastmanconf.yaml"
	content, err := sampledata.GetSampleContent(configFileName)
	if err != nil {
		return err
	}

	configFilePath := filepath.Join(currentDir, configFileName)

	exists, err := FileOrDirExists(configFilePath)
	if err != nil {
		return err
	}

	err = CreateFile(configFilePath, []byte(content))
	if err != nil {
		return err
	}

	if exists {
		fmt.Printf("fastman reset in: %s\n", currentDir)
	} else {
		fmt.Printf("fastman initialized in: %s\n", currentDir)
	}

	return nil
}
