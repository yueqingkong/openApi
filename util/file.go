package util

import "os"

func MkDir(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}
