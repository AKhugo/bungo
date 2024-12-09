package fi

import (
	"os"
	"fmt"
)

// FileExists check if a file exist
// if the file does not exist, it returns false
// if there is an error while checking the file, it returns false
func FileExists(path string) bool {

	if path == "" {
		return false
	}

	_, err := os.Stat(path)
	
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	return true

}

// GetFileData read a file and return its content as a byte slice.
// If the file does not exist, it returns an error.
// If there is an error while reading the file, it returns an error.
func GetFileData(path string) ([]byte, error) {


	
	if !FileExists(path) {
		return nil, fmt.Errorf("file %s does not exist: ", path)
	}


	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", path, err.Error())
	}

	

	return data, nil

}


