package cmd

import (
	"fmt"
	"os"
	"github.com/koffihuguesagossadou/bungo/pkg/fi"
	"github.com/koffihuguesagossadou/bungo/pkg/format"
	"github.com/spf13/cobra"
)

// encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a file in base64",
	Long:  `Encode a file in base64`,
	RunE: func (cmd *cobra.Command, args []string) error {
		
		// Récupérer les flags
		inputFile, err := cmd.Flags().GetString("i")

		if err != nil {
			return err
		}



		if inputFile == "" {
			return fmt.Errorf("error encoding file: input file is missing")
		}

		data, err := fi.GetFileData(inputFile)

		if err != nil {
			return fmt.Errorf( "error while reading file %s: %s", inputFile, err.Error())
		}

		encoded, err := format.EncodeToBase64(data)

		if err != nil {
			
			return fmt.Errorf( "error encoding file %s: %s", inputFile, err.Error())
		}

		outputFile, err := cmd.Flags().GetString("o")

		if err != nil {
			return err
		}


		// write encoded file
		if outputFile == "" {	
			return fmt.Errorf("error encoding file: output file is missing")
		}

		// check if output file exist
		outputFileExists := fi.FileExists(outputFile)

		// // if not exist create it
		if !outputFileExists {
			_, err := os.Create(outputFile)
			if err != nil {
				
				return fmt.Errorf("error encoding file : error creating output file %s", outputFile)
			}
		}

		// now write encoded file
		
		os.WriteFile(outputFile, []byte(encoded), 0644)

		// message to user
		fmt.Println("file encoded in base64. Please check your output file : " + outputFile)

		return nil
	},


}