package operation

import (
	"fmt"

	"github.com/koffihuguesagossadou/bungo/common"
	"github.com/koffihuguesagossadou/bungo/pkg/utils"
	"github.com/spf13/cobra"
)


func Compress( cmd *cobra.Command, args []string) error {


    inputFile, err := cmd.Flags().GetString("i")

    if err != nil {
        return utils.ThrowError(err, common.ERROR_COMPRESSING_FILE);
    }


    if inputFile == "" {
        return fmt.Errorf("error encoding file: input file is missing")
    }
    


    err = utils.ImageCompressor(inputFile, nil)

    if err != nil {
        return utils.ThrowError(err, common.ERROR_COMPRESSING_FILE);
    }

    return nil;
}
