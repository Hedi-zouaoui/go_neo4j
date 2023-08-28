/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mlmcli",
	Short: "CLI for MLM tree manipulation",
	Long: `mlmcli is a command-line interface (CLI) tool designed to facilitate the manipulation of MLM (Multi-Level Marketing) trees. 
It provides various commands to perform operations on MLM tree structures, including adding users, deleting users, and transferring users between nodes.

The mlmcli tool simplifies the process of interacting with MLM trees by providing a command-based interface that can be easily integrated into scripts or used directly from the command line. 
You can leverage its functionality to automate MLM tree management tasks, integrate with other systems, or perform ad hoc operations.

Each command supports a set of options and arguments to customize the operation. Use the '--help' flag with any command to display more detailed usage information and available options.

For more information and examples, please refer to the documentation or visit the project repository.

Enjoy using mlmcli for MLM tree manipulation!
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mlmcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
