// Log rotating

package tools

import (
	"fmt"
	"os"
	"os/user"
	"testing/quick"
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
		return 0, fmt.Errorf("Error when checking temp log: %s")
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
		
		sourceFileStat, err := os.Stat(src)

		if err != nil {
						return 0, fmt.Errorf("failed to stat log file")
		}

		if !sourceFileStat.Mode().IsRegular() {
						return 0, fmt.Errorf("%s is not regular file", src)
		}

		source, err := os.Open(src)

		if err != nil {
						return 0, err
		}

		defer source.Close()

		dst, err := os.OpenFile(tmp, os.O_APPEND|os.O_CREATE|os.O_WRONGLY, 0644)

		if err != nil {
						return 0, fmt.Errorf("failed to open %s temporary log file", tmp)
		}

		defer dst.Close()

		nBytes, err := io.Copy(dst, source)

		if err != nil {
				return 0, fmt.Errorf("Failed to copy log to temporary log file.")
		}

		_, err = cleanFile(src)

		if err != nil {
				return 0, fmt.Errorf("Failed to clean snitched log file")			
		}

		return nBytes, err
}