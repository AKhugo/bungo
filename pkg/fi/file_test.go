package fi

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {

	// Test case 1: oath empty

	if FileExists("") {
		t.Errorf("FileExists should return false for an empty path")
	}

	// Test case 2: File does not exist
	if FileExists("file-does-not-exist.img") {
		t.Errorf("FileExists(\"non_existent_file.txt\") = true; want false")
	}

	// Test case 3: File exists
	existingFile := "file-exists.txt"
	err := os.WriteFile(existingFile, []byte("Hello, world!"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(existingFile)

	if !FileExists(existingFile) {
		t.Errorf("FileExists(\"%s\") = false; want true", existingFile)
	}
}

func TestGetFileData(t *testing.T) {
	// Cas 1 : Fichier inexistant
	_, err := GetFileData("non_existent_file.txt")
	if err == nil {
		t.Error("GetFileData(\"non_existent_file.txt\") = nil; want error")
	}

	// Cas 2 : Fichier texte existant
	textFile := "test_file.txt"
	textContent := "Hello, World!"
	err = os.WriteFile(textFile, []byte(textContent), 0644)
	if err != nil {
		t.Fatalf("Erreur lors de la création du fichier : %v", err)
	}
	defer os.Remove(textFile) // Nettoyage après le test

	data, err := GetFileData(textFile)
	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error", textFile, err)
	}
	if string(data) != textContent {
		t.Errorf("GetFileData(\"%s\") = \"%s\"; want \"%s\"", textFile, string(data), textContent)
	}

	// Cas 3 : Fichier binaire existant
	binaryFile := "test_image.png"
	binaryContent := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A} // En-tête PNG
	err = os.WriteFile(binaryFile, binaryContent, 0644)
	if err != nil {
		t.Fatalf("Erreur lors de la création du fichier : %v", err)
	}
	defer os.Remove(binaryFile) // Nettoyage après le test

	binaryData, err := GetFileData(binaryFile)
	if err != nil {
		t.Errorf("GetFileData(\"%s\") = error %v; want no error", binaryFile, err)
	}
	if len(binaryData) != len(binaryContent) {
		t.Errorf("GetFileData(\"%s\"): taille = %d; want %d", binaryFile, len(binaryData), len(binaryContent))
	}
	for i := range binaryContent {
		if binaryData[i] != binaryContent[i] {
			t.Errorf("GetFileData(\"%s\"): byte[%d] = %x; want %x", binaryFile, i, binaryData[i], binaryContent[i])
		}
	}
}


