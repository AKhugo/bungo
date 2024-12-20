package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/koffihuguesagossadou/bungo/pkg/command"
	"github.com/spf13/cobra"
)



func TestEncoder(t *testing.T) {

	// Case 1 : input file is missing
	cmd := &cobra.Command{}
	cmd.Flags().StringP("i", "i", "", "input file") // Simule un flag sans valeur
	cmd.Flags().StringP("o", "o", "output.txt", "output file")

	err := command.EncodeCmd.RunE(cmd, []string{})

	if( err != nil && !strings.Contains(err.Error(), "error encoding file") ) {
		t.Errorf("encodeCmd.RunE() = nil, want error 'input file is missing'")
	}

	// Case 2 : input file does not exist
	cmd.Flags().Set("i", "non_existent_file.txt")
	err = command.EncodeCmd.RunE(cmd, []string{})
	if err == nil || !strings.Contains(err.Error(), "error while reading file") {
		t.Errorf("encodeCmd.RunE() = nil, want error 'file does not exist'")
	}


	inputFile, err := filepath.Abs("../test/files/good-docx-test.docx");

	if err != nil {
		t.Errorf("FileExists(\"%s\") = false; want true", inputFile)
	}

	outputFile := "encoded_file.txt"

	// Test case 1: Encode a file
	command.EncodeCmd.Flags().StringP("i", "i", inputFile, "input file")
	command.EncodeCmd.Flags().StringP("o", "o", outputFile, "output file")

	// Delete the output file if it exists
	os.Remove(outputFile)


	if err := command.EncodeCmd.Execute(); err != nil {
		t.Errorf("EncodeCmd.Execute() error = %v, want nil", err)
	}


	// delete output file
	os.Remove(outputFile)

	// Cas 4 : No output file provided
	cmd.Flags().Set("o", "")
	err = command.EncodeCmd.RunE(cmd, []string{})
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
	err = command.EncodeCmd.RunE(cmd, []string{})
	if err == nil || strings.Contains(err.Error(), "error creating output file") {
		t.Errorf("encodeCmd.RunE() = nil, want error 'error creating output file': %v", err)
	}

	os.Remove(f.Name())
	

}


func TestDecoder(t *testing.T) {

	// Case 1 : input file is missing
	cmd := &cobra.Command{}
	cmd.Flags().StringP("i", "i", "", "input file") // Simule un flag sans valeur
	cmd.Flags().StringP("o", "o", "output.txt", "output file")


	err := command.DecodeCmd.RunE(cmd, []string{})
	if err != nil && err == cmd.MarkFlagRequired("i") {
		t.Errorf("decodeCmd.RunE() = nil, want error 'input file is missing'")
	}

	// Case 2 : input file does not exist
	cmd.Flags().Set("i", "non_existent_file.txt")
	err = command.DecodeCmd.RunE(cmd, []string{})
	if err == nil && strings.Contains(err.Error(), "does not exist") {
		t.Errorf("decodeCmd.RunE() = nil, want error 'file does not exist' %v", err)
	}

	// Cas 3 : invalid input file

	inputFile, err := filepath.Abs("../test/files/input-base64.txt");

	if err != nil {
		t.Errorf("FileExists(\"%s\") = false; want true", inputFile)
	}
	

	outputFile := "decoded_file"

	// Test case 1: Decode a file
	command.DecodeCmd.Flags().StringP("i", "i", inputFile, "input file")
	command.DecodeCmd.Flags().StringP("o", "o", outputFile, "output file")

	// Delete the output file if it exists
	os.Remove(outputFile)

	if err := command.DecodeCmd.Execute(); err != nil {
		t.Errorf("DecodeCmd.Execute() error = %v, want nil", err)
	}

	// delete output file
	os.Remove(outputFile)

	// Case 4 : No output file provided
	cmd.Flags().Set("o", "")
	err = command.DecodeCmd.RunE(cmd, []string{})
	if err == nil && strings.Contains(err.Error(), "output file is missing"){
		t.Errorf("decodeCmd.RunE() = nil, want error 'output file is missing'")
	}

}