package command

import (
	"github.com/koffihuguesagossadou/bungo/pkg/operation"
	"github.com/spf13/cobra"
)

// encode command
var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a file in base64",
	Long:  `Encode any file in base64`,
	RunE: func (cmd *cobra.Command, args []string) error{
		
		return operation.Encoder(cmd, args)
	},
}

// decode command

var DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a file in base64",
	Long:  `Decode any file in base64`,
	RunE: func (cmd *cobra.Command, args []string) error{
		
		return operation.Decoder(cmd, args)
	},
}