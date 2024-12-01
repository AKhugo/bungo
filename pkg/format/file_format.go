package format

import (
	"fmt"
	"encoding/base64"

)

// decode file to base64
func EncodeToBase64(file []byte) (string, error) {

	if file == nil {
		return "", fmt.Errorf("file is empty")
	}

	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded, nil
}

