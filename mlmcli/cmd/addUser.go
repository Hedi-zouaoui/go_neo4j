package cmd

import (
	"fmt"
	//"os"

	"github.com/mlmcli/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user to the MLM tree",
Long: "The add command allows you to create a new user under a specific node in the MLM tree. Specify the necessary details such as the user's name and the target node where the user will be added.",
	Run: func(cmd *cobra.Command, args []string) {

			to, _ := cmd.Flags().GetInt64("to")
		name, _ := cmd.Flags().GetString("name")
	
		fmt.Println("value of to " , to )
		req := utils.AddUser_request{
			To:   to,
			Name: name,
		}
		fmt.Printf("Creating task %+v\n", req)

		// Storing task in file, you can call REST APIs also
		resp := utils.AddUser(req)
		fmt.Println(resp)
		
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Int64P("to", "t", 0, "the id of the node you want to add to")
	addCmd.Flags().StringP("name", "d", "", "the name of the new user ")
}