package route

import (
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createRoutesDir(d *config.RouteProps) error {
	_, err := os.Stat(d.RouteSubDir)

	if os.IsNotExist(err) {
		os.MkdirAll(d.RouteSubDir, os.ModePerm)
	}

	return nil
}
