package test

import (
	"path/filepath"
	"testing"

	"github.com/koffihuguesagossadou/bungo/pkg/fi"
)

func TestFileExists(t *testing.T) {

	// Test case 1: path empty

	if fi.FileExists("") {
		t.Errorf("FileExists should return false for an empty path")
	}

	// Test case 2: File does not exist
	if fi.FileExists("file-does-not-exist.img") {
		t.Errorf("FileExists(\"non_existent_file.txt\") = true; want false")
	}

	// Test case 3: File exists

	// get file in test/files directory
	existingFile, err := filepath.Abs("../test/files/good-docx-test.docx");

	if err != nil {
		t.Errorf("FileExists(\"%s\") = false; want true", existingFile)
	}


	if !fi.FileExists(existingFile) {
		t.Errorf("FileExists(\"%s\") = false; want true", existingFile)
	}
}

// TestGetFileData tests the GetFileData function.
//
// It tests the following cases:
//
// 1. A non-existent file.
// 2. A text file with known content.
// 3. A binary file with known content.
func TestGetFileData(t *testing.T) {
	// Cas 1 : Fichier inexistant
	_, err := fi.GetFileData("non_existent_file.txt")
	if err == nil {
		t.Error("GetFileData(\"non_existent_file.txt\") = nil; want error")
	}

	// Cas 2 : Fichier texte existant (txt)
	textFile, err := filepath.Abs("../test/files/good-txt-test.txt");
	
	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error, Error getting file path:", textFile, err)
	}

	data , err := fi.GetFileData(textFile)
	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error", textFile, err)
	}

	if len(data) == 0 {
		t.Errorf("GetFileData(\"%s\"): data = %q; want non-empty string", textFile, data)
	}


	// Cas 3 : Fichier binaire existant
	binaryFile, err := filepath.Abs("../test/files/good-jpeg-test.jpeg");
	
	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error, Error getting file path:", binaryFile, err)
	}

	binaryData, err := fi.GetFileData(binaryFile)

	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error", binaryFile, err)
	}

	if len(binaryData) == 0 {
		t.Errorf("GetFileData(\"%s\"): data = %q; want non-empty string", binaryFile, binaryData)
	}

}


