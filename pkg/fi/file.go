package fi

import (
	"os"
	"fmt"
)

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


