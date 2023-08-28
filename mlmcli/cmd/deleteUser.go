package cmd

import (
	"fmt"
	//"os"

	"github.com/mlmcli/utils"
	"github.com/spf13/cobra"
)
/* 
type DeleteUser_request struct {
  Id int64  `json:"id"`

}

}*/ 

// createCmd represents the create command
var deleteCmd = &cobra.Command{
Use: "delete",
Short: "Delete a user from the MLM tree",
Long: "The delete command allows you to remove a user from the MLM tree. Specify the user to be deleted, and the command will remove them from the tree structure.",

	Run: func(cmd *cobra.Command, args []string) {

			id, _ := cmd.Flags().GetInt64("id")
	
	
		fmt.Println("value of to " , id )
		req := utils.DeleteUser_request  {
			Id:   id,
		
		}
		fmt.Printf("Creating task %+v\n", req)

		// Storing task in file, you can call REST APIs also
		resp := utils.DeleteUser(req)
		fmt.Println(resp)
		
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd )
	deleteCmd .Flags().Int64P("id", "i", 0, "the id of the node you want to delete")

}