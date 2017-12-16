package commands

import (
	"github.com/spf13/cobra"
	"github.com/kj187/aws-helper/commands/ec2"
)

var tags []string
var filters []string
var columns []string
var removeColumns []string

var listCommand = &cobra.Command{
	Use: "ec2:list",
	Short: "List all EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		ec2.ListInstances(region, tags, filters, columns, removeColumns)
	},
}

func init() {
	listCommand.Flags().StringSliceVarP(&tags, "tag", "t", nil, "filter with tag (Example: Name:Jenkins)")
	listCommand.Flags().StringSliceVarP(&filters, "filter", "f", nil, "filter with column (Example: InstanceType:t2.micro)")
	listCommand.Flags().StringSliceVarP(&columns, "column", "c", nil, "add additional column (tag)")
	listCommand.Flags().StringSliceVarP(&removeColumns, "remove-column", "C", nil, "remove default column")
	RootCmd.AddCommand(listCommand)
}
