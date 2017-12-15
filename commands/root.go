package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "aws-inspector",
	Short: "AWS Inspector",
	Run: func(cmd *cobra.Command, args []string) {},
}