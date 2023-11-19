package service

import (
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createServiceDir(d *config.ServiceProps) error {

	_, err := os.Stat(d.ServiceSubDir)

	if os.IsNotExist(err) {
		os.MkdirAll(d.ServiceSubDir, os.ModePerm)
	}

	return nil
}

func createServiceTestDir(d *config.ServiceProps) error {

	_, err := os.Stat(d.TestSubDir)

	if os.IsNotExist(err) {
		os.MkdirAll(d.TestSubDir, os.ModePerm)
	}

	return nil
}
