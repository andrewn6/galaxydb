// Log rotating

package tools

import (
	"os"
)

func LogMustRotate(logfile string, maxSize int64) (bool, error) {
	fi, err := os.Stat(logfile)

	if err != nil {
		return false, err

	}

	// get the log file size
	size := fi.Size()

	if size >= maxSize {
		return true, nil
	}
	return false, nil
}
