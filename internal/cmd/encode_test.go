package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/koffihuguesagossadou/bungo/pkg/fi"
	"github.com/spf13/cobra"
)

func TestDo(t *testing.T) {

	// Case 1 : input file is missing
	cmd := &cobra.Command{}
	cmd.Flags().StringP("i", "i", "", "input file") // Simule un flag sans valeur
	cmd.Flags().StringP("o", "o", "output.txt", "output file")

	err := encodeCmd.RunE(cmd, []string{})
	if err == nil || err.Error() != "error encoding file: input file is missing" {
		t.Errorf("encodeCmd.RunE() = nil, want error 'input file is missing'")
	}

	// Case 2 : input file does not exist
	cmd.Flags().Set("i", "non_existent_file.txt")
	err = encodeCmd.RunE(cmd, []string{})
	if err == nil || !strings.Contains(err.Error(), "error while reading file") {
		t.Errorf("encodeCmd.RunE() = nil, want error 'file does not exist'")
	}


	inputFile, err := filepath.Abs("../../test/files/good-docx-test.docx");

	if err != nil {
		t.Errorf("FileExists(\"%s\") = false; want true", inputFile)
	}

	outputFile := "encoded_file.txt"

	// Test case 1: Encode a file
	encodeCmd.Flags().StringP("i", "i", inputFile, "input file")
	encodeCmd.Flags().StringP("o", "o", outputFile, "output file")

	// Delete the output file if it exists
	os.Remove(outputFile)


	if err := encodeCmd.Execute(); err != nil {
		t.Errorf("EncodeCmd.Execute() error = %v, want nil", err)
	}

	// check if output file have been created
	if !fi.FileExists(outputFile) {
		t.Errorf("FileExists(\"%s\") = false; want true", outputFile)
	}

	// delete output file
	os.Remove(outputFile)

	// Cas 4 : No output file provided
	cmd.Flags().Set("o", "")
	err = encodeCmd.RunE(cmd, []string{})
	if err == nil || strings.Contains(err.Error(), "output file is missing"){
		t.Errorf("encodeCmd.RunE() = nil, want error 'output file is missing'")
	}

	// Cas 5 : no way to write in file (permissions)

	// create a file with no write permissions
	f, err := os.OpenFile("nonwritable.txt", os.O_CREATE|os.O_RDONLY, 0444)
	if err != nil {
		t.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(inputFile)

	cmd.Flags().Set("o", f.Name())
	err = encodeCmd.RunE(cmd, []string{})
	if err == nil || strings.Contains(err.Error(), "error creating output file") {
		t.Errorf("encodeCmd.RunE() = nil, want error 'error creating output file': %v", err)
	}

	os.Remove(f.Name())
	

}