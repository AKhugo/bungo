package command

import (
	"fmt"
	"os"
	"github.com/koffihuguesagossadou/bungo/common"
	"github.com/koffihuguesagossadou/bungo/pkg/fi"
	"github.com/koffihuguesagossadou/bungo/pkg/format"
	"github.com/koffihuguesagossadou/bungo/pkg/utils"
	"github.com/spf13/cobra"
)

// encode command
var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a file in base64",
	Long:  `Encode any file in base64`,
	RunE: func (cmd *cobra.Command, args []string) error {
		
		// Récupérer les flags
		inputFile, err := cmd.Flags().GetString("i")

		if err != nil {
			return fmt.Errorf("error encoding file: %s", err.Error());
		}


		if inputFile == "" {
			return fmt.Errorf("error encoding file: input file is missing")
		}

		data, err := fi.GetFileData(inputFile)

		if err != nil {
			return fmt.Errorf( "error while reading file %s: %s", inputFile, err.Error())
		}

		encoded, err := format.EncodeToBase64(data);

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
		err = os.WriteFile(outputFile, []byte(encoded), 0644)

		if err != nil {
			return fmt.Errorf("error encoding file : error writing output file %s", outputFile)
		}

		// message to user
		fmt.Println("file encoded in base64. Please check your output file : " + outputFile)


		return nil
	},


}

// decode command

var DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a file in base64",
	Long:  `Decode any file in base64`,
	RunE: func (cmd *cobra.Command, args []string) error {
		  

		// get flags 

		inputFile, err := cmd.Flags().GetString("i")

		if err != nil {
			return utils.ThrowError(err, common.ERROR_DECODING_FILE)
		}


		if inputFile == "" {

			err = cmd.MarkFlagRequired("i")

			return utils.ThrowError(err, common.ERROR_DECODING_FILE)
		}


		// read file
		data, err := fi.GetFileData(inputFile);

		if err != nil {

			return utils.ThrowError(err, "reading input file");
		}


		// decode file
		decoded, err := format.DecodeBase64(string(data));

		if err != nil {

			return utils.ThrowError(err, common.ERROR_DECODING_FILE);
		}


		// write decoded file
		outputFile, err := cmd.Flags().GetString("o")

		if err != nil {
			return utils.ThrowError(err, common.ERROR_DECODING_FILE)
		}


		// check if output file exist
		outputFileExists := fi.FileExists(outputFile);

		// // if not exist create it
		if !outputFileExists {
			_, err := os.Create(outputFile)
			if err != nil {
				return utils.ThrowError(err, common.ERROR_DECODING_FILE)
			}
		}

		// now write decoded file
		err = os.WriteFile(outputFile, decoded, 0644)

		if err != nil {
			return utils.ThrowError(err, common.ERROR_DECODING_FILE)
		}

		// message to user
		fmt.Println("file decoded in base64. Please check your output file : " + outputFile)
		return nil

	},
}