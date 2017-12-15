package commands

import (
	"github.com/spf13/cobra"
	"github.com/kj187/aws-inspector/commands/ec2"
)

var listCommand = &cobra.Command{
	Use: "ec2:list",
	Short: "List all EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		ec2.ListInstances(region)
	},
}

func init() {
	RootCmd.AddCommand(listCommand)
}