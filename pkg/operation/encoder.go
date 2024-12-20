package operation

import (
	"fmt"
	"os"
	"net/http"
	"github.com/koffihuguesagossadou/bungo/common"
	"github.com/koffihuguesagossadou/bungo/pkg/utils"
	"github.com/spf13/cobra"
)



func Encoder(cmd *cobra.Command, args []string) error {
		
	// Récupérer les flags
	inputFile, err := cmd.Flags().GetString("i")

	if err != nil {
		return fmt.Errorf("error encoding file: %s", err.Error());
	}


	if inputFile == "" {
		return fmt.Errorf("error encoding file: input file is missing")
	}

	data, err := utils.GetFileData(inputFile)

	if err != nil {
		return fmt.Errorf( "error while reading file %s: %s", inputFile, err.Error())
	}

	encoded, err := utils.EncodeToBase64(data);

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
	outputFileExists := utils.FileExists(outputFile)

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
}


func Decoder(cmd *cobra.Command, args []string) error {
		  

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
	data, err := utils.GetFileData(inputFile);

	if err != nil {

		return utils.ThrowError(err, "reading input file");
	}


	// decode file
	decoded, err := utils.DecodeBase64(string(data));

	if err != nil {

		return utils.ThrowError(err, common.ERROR_DECODING_FILE);
	}


	//detect MIME type
	mimetype := http.DetectContentType(decoded);


	// get output file
	outputFile, err := cmd.Flags().GetString("o")

	if err != nil {
		return utils.ThrowError(err, common.ERROR_DECODING_FILE)
	}

	// check if output file exist
	outputFileExists := utils.FileExists(outputFile);

	// // if not exist create it
	if !outputFileExists {
		_ , err := os.Create(outputFile)

		if err != nil {
			return utils.ThrowError(err, common.ERROR_DECODING_FILE)
		}
	}

	// empty file
	if err := os.Truncate(outputFile, 0); err != nil {
		return utils.ThrowError(err, common.ERROR_DECODING_FILE)
	}


	os.Remove(outputFile)

	// attribute file extension
	if extension, ok := common.ExtMap[mimetype]; ok {
		outputFile = outputFile + extension
	}

	

	fmt.Print("output file : " + outputFile)

	
	// now write decoded file
	err = os.WriteFile(outputFile, decoded, 0644)

	if err != nil {
		return utils.ThrowError(err, common.ERROR_DECODING_FILE)
	}

	// message to user
	fmt.Println("file decoded in base64. Please check your output file : " + outputFile)
	return nil

}