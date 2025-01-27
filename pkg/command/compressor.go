package command

import (
	"github.com/koffihuguesagossadou/bungo/pkg/operation"
	"github.com/spf13/cobra"
)

var CompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress image",
	Long:  `Compress any image`,
	RunE: func (cmd *cobra.Command, args []string) error{
		
		return operation.Compress(cmd, args)
	},
}