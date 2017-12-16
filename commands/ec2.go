package commands

import (
	"github.com/spf13/cobra"
	"github.com/kj187/aws-helper/commands/ec2"
)

var tags string

var listCommand = &cobra.Command{
	Use: "ec2:list",
	Short: "List all EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		ec2.ListInstances(region, tags)
	},
}

func init() {
	listCommand.Flags().StringVarP(&tags, "tags", "t", "", "Filter with tags. Example: \"Name:Jenkins, Environment:Prod\")")
	RootCmd.AddCommand(listCommand)
}
