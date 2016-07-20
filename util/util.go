package util

import (
	"os"
)

// DcosDomain is the domain that identifies a request as one that should
// be processed.
const DcosDomain string = ".mydcos.directory"

// MaxPortLength is the maximum number of digits a (network) port can be.
const MaxPortLength int = 5

// RmIfExist removes a file (given the path) if it exists.
func RmIfExist(path string) error {
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}
