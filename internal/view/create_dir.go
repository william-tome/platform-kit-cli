package view

import (
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createViewDir(d *config.ViewProps) error {
	_, err := os.Stat(d.ViewSubDir)

	if os.IsNotExist(err) {
		os.MkdirAll(d.ViewSubDir, os.ModePerm)
	}

	return nil
}

func createBlockDir(d *config.ViewProps) error {
	_, err := os.Stat(d.BlockSubDir)

	if os.IsNotExist(err) {
		os.MkdirAll(d.BlockSubDir, os.ModePerm)
	}

	return nil
}
