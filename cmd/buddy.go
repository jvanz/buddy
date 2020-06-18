package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	currentDir string

	buddyCmd = &cobra.Command{
		Use:              "buddy",
		Short:            "Utilitary tool to get info from supportutils",
		TraverseChildren: true,
	}
)

func init() {
	buddyCmd.AddCommand(getCmd)
	buddyCmd.PersistentFlags().StringVarP(&currentDir, "dir", "d", "", "Directory with the supportutil files")
}

func Execute() {
	if err := buddyCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getSupportutilsDir() string {
	if len(currentDir) > 0 {
		return currentDir
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
