package commands

import (
	"github.com/spf13/cobra"
	"github.com/kj187/aws-inspector/commands/ec2"
)

var filter string

var listCommand = &cobra.Command{
	Use: "ec2:list",
	Short: "List all EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		ec2.ListInstances(region, filter)
	},
}

func init() {
	listCommand.Flags().StringVarP(&filter, "filter", "f", "", "Set a filter via tags. Example: \"Name:Jenkins, Environment:Prod\")")
	RootCmd.AddCommand(listCommand)
}
