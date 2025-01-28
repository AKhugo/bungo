package operation

import (
	"fmt"

	"github.com/koffihuguesagossadou/bungo/common"
	"github.com/koffihuguesagossadou/bungo/pkg/utils"
	"github.com/spf13/cobra"
)


func Compress( cmd *cobra.Command, args []string) error {


    inputFile, err := cmd.Flags().GetString("input")

    if err != nil {
        return utils.ThrowError(err, common.ERROR_COMPRESSING_FILE);
    }



    if inputFile == "" {
        return fmt.Errorf("error encoding file: input file is missing")
    }

    quality, err := cmd.Flags().GetInt("quality")

    if err != nil {
        return utils.ThrowError(err, common.ERROR_COMPRESSING_FILE);
    }



    err = utils.ImageCompressor(inputFile, &quality)

    if err != nil {
        return utils.ThrowError(err, common.ERROR_COMPRESSING_FILE);
    }

    return nil;
}
