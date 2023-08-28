package cmd

import (
	"fmt"
	//"os"

	"github.com/mlmcli/utils"
	"github.com/spf13/cobra"
)
/* type TransferUser_request struct {
	  To int64  `json:"to"`
	  From int64  `json:"from"`
     Methode string  `json:"methode"`

}*/ 

// createCmd represents the create command
var transferCmd = &cobra.Command{
Use: "transfer",
Short: "Transfer a user to another node in the MLM tree",
Long: "The transfer command allows you to transfer a user from one node to another in the MLM tree. Specify the user to be transferred and the target node where the user should be moved.",

	Run: func(cmd *cobra.Command, args []string) {

			to, _ := cmd.Flags().GetInt64("to")
		from, _ := cmd.Flags().GetInt64("from")
		methode , _ := cmd.Flags().GetString("methode")
	
		fmt.Println("value of to " , to )
		req := utils.TransferUser_request {
			To:   to,
			From: from,
			Methode : methode  , 
		}
		fmt.Printf("Creating task %+v\n", req)

		// Storing task in file, you can call REST APIs also
		resp := utils.TransferUser(req)
		fmt.Println(resp)
		
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	transferCmd.Flags().Int64P("to", "t", 0, "the id of the node you want to move to")
	transferCmd.Flags().Int64P("from", "f", 0 , "the id of the node you want to move ")
	transferCmd.Flags().StringP("methode", "w", "with", "with its kids or without , the default is with ")

}