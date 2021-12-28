// Log rotating

package tools

import (
	"fmt"
	"os"
	"os/user"
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

func Rotate(filename string, tmpFilename string) (int64, error) {
	usr, _ := user.Current()
	configPath := usr.HomeDir

	src := configPath + filename
	tmp := configPath + tmpFilename

	exists, err := tmpFileExists(tmp)

	if err != nil {
		return 0, fmt.Errorf("Erorr when checking temp log: %s", err.Error())
	}

	if exists {
		_, err := cleanFile(tmp)

		if err != nil {
			return 0, fmt.Errorf("failed to clean temporary log")
		}

		// check if log exists, then clear it
		if exists {
			_, err := cleanFile(tmp)

			if err != nil {
				return 0, fmt.Errorf("failed to clean temp log")
			}
		}
	}

}
