package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/koffihuguesagossadou/bungo/common"
)

// decode file to base64
func EncodeToBase64(file []byte) (string, error) {

	if file == nil {
		return "", ThrowError(fmt.Errorf("file is empty"), "encoding file")
	}

	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded, nil
}

func DecodeBase64(encoded string) ([]byte, error) {

	if encoded == "" {

		err := fmt.Errorf("encoded is empty")

		return nil, ThrowError(err, "decoding file")
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {

		return nil, ThrowError(err, "decoding file")

	}

	return decoded, nil
}

func ImageCompressor(inputPath string, quality *int) error {

	// if quality is not provided, set it to 70
	if quality == nil {
		defaultQuality := 70
		quality = &defaultQuality
	}

	if *quality < 1 || *quality > 100 {
		return ThrowError(fmt.Errorf("quality is out of range"), "compressing file")
	}

	// Open the input file
	input, err := os.Open(inputPath)
	if err != nil {
		return ThrowError(err, "compressing file")
	}
	defer input.Close()

	// decode the image
	image, format, err := image.Decode(input)
	if err != nil {
		return ThrowError(err, "compressing file")
	}

	// get current path
	currentPath, err := os.Getwd()

	if err != nil {
		return ThrowError(err, common.ERROR_COMPRESSING_FILE)
	}

	// get file name without extension
	fileName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))

	outputPath := currentPath + "/" + fileName + "_compressed." + format

	// output image
	output, err := os.Create(outputPath)
	if err != nil {
		return ThrowError(err, "compressing file")
	}
	defer output.Close()

	opts := jpeg.Options{Quality: *quality}

	switch format {
	case "jpeg":
		jpeg.Encode(output, image, &opts)
	case "png":
		png.Encode(output, image)
	case "jpg":
		jpeg.Encode(output, image, &opts)
	default:
		return ThrowError(fmt.Errorf("unsupported format: %s", format), "compressing file")
	}

	return nil

}

// func FileCompressor(file []byte) ([]byte, error) {

// }
