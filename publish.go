package surge

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pierrre/archivefile/zip"
)

// Publish the project to the domain
func (s *Client) Publish(projectPath, domain string) error {
	fmt.Println(fmt.Sprintf("publish %s -> %s", projectPath, domain))

	fullProjectPath, err := filepath.Abs(projectPath)
	if err != nil {
		return err
	}

	fmt.Println(fullProjectPath)

	zipPath, err := filepath.Abs("./publish.zip")
	if err != nil {
		return err
	}

	outputFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}

	zip.Archive(fullProjectPath, outputFile, func(archivePath string) {
		fmt.Println(archivePath)
	})

	return nil
}
