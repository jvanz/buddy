package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get info from supportutils",
}

func init() {
	getCmd.AddCommand(getProcess)
}
