// +build !linux

package lumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo, _, _ int) error {
	return nil
}
