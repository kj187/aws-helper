package commands

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var region string
var profile string

const DefaultRegion = "eu-central-1"

var RootCmd = &cobra.Command{
	Use: "aws-helper",
	Short: "AWS CLI Helper",
	Long: "AWS Helper is a lightweight CLI tool to work with AWS.\nAuthor: Julian Kleinhans <mail@kj187.de>, alias @kj187",
	Example: "aws-helper ec2:list -c Name -C KeyName -f AZ:eu-central-1",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&region, "region", "r", DefaultRegion, "set region")
	RootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "set aws profile")
}

func Execute() {
	/*
	https://docs.aws.amazon.com/de_de/sdk-for-go/v1/developer-guide/configuring-sdk.html

	os.Setenv("AWS_PROFILE", test-account)

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		color.Red("No AWS_ACCESS_KEY_ID env var available!")
		return
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		color.Red("No AWS_SECRET_ACCESS_KEY env var available!")
		return
	}*/

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "aws-helper error: %s\n", err)
		RootCmd.Usage()
		os.Exit(1)
	}
}