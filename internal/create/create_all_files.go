package create

import (
	"github.com/william-tome/platform-kit-cli/internal/route"
	"github.com/william-tome/platform-kit-cli/internal/service"
	"github.com/william-tome/platform-kit-cli/internal/view"
)

func CreateAllFiles(name string) error {
	if err := service.CreateService(name); err != nil {
		return err
	}

	if err := route.CreateRoute(name, name); err != nil {
		return err
	}

	if err := view.CreateView(name, name, name); err != nil {
		return err
	}

	return nil
}
