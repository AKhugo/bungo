package cmd

import (
	"testing"
)

func TestDo(t *testing.T) {

	


	// Test case 1: Encode a file
	encodeCmd.Flags().StringP("i", "i", "test_file.txt", "input file")
	encodeCmd.Flags().StringP("o", "o", "encoded_file.txt", "output file")

	if err := encodeCmd.Execute(); err != nil {
		t.Errorf("EncodeCmd.Execute() error = %v, want nil", err)
	}

	// Test case 2: Encode a non-existent file
	encodeCmd.Flags().StringP("i", "i", "non_existent_file.txt", "input file")
	encodeCmd.Flags().StringP("o", "o", "encoded_file.txt", "output file")

	if err := encodeCmd.Execute(); err == nil {
		t.Errorf("EncodeCmd.Execute() error = nil, want error")
	}

}