package format

import (
	"encoding/base64"
	"fmt"

	"github.com/koffihuguesagossadou/bungo/pkg/utils"
)

// decode file to base64
func EncodeToBase64(file []byte) (string, error) {

	if file == nil {
		return "", utils.ThrowError(fmt.Errorf("file is empty"), "encoding file")
	}

	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded, nil
}


func DecodeBase64(encoded string) ([]byte, error) {

	if encoded == "" {

		err := fmt.Errorf("encoded is empty")

		return nil, utils.ThrowError( err, "decoding file")
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {

		return nil, utils.ThrowError(err, "decoding file")

	}

	return decoded, nil
}

