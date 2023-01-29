package shared

import (
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var ErrNetworkNotSupported = errors.New("network not supported")

func PrintError(err error) {
	color.HiRed("ERROR: %s", err.Error())
}
