package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var region string
var profile string
var accessKey string
var secretKey string

// DefaultRegion is the default region
const DefaultRegion = "eu-central-1"

// RootCommand is the root command
var RootCommand = &cobra.Command{
	Use:     "aws-helper",
	Short:   "AWS CLI Helper",
	Long:    "The AWS Helper is a go based command line interface utility for AWS.\nAuthor: Julian Kleinhans <mail@kj187.de>, alias @kj187",
	Example: "aws-helper ec2:list -c Name -C KeyName -f AZ:eu-central-1",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCommand.PersistentFlags().StringVarP(&region, "region", "r", DefaultRegion, "set region")
	RootCommand.PersistentFlags().StringVarP(&profile, "profile", "p", "", "set aws profile")
	RootCommand.PersistentFlags().StringVarP(&accessKey, "access_key", "a", "", "set aws access_key")
	RootCommand.PersistentFlags().StringVarP(&secretKey, "secret_key", "s", "", "set aws secret_key")
}

// Execute the root command
func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "aws-helper error: %s\n", err)
		RootCommand.Usage()
		os.Exit(1)
	}
}

func loadAwsCredentials() {
	if profile == "" {
		if accessKey == "" || secretKey == "" {
			if os.Getenv("AWS_PROFILE") == "" {
				if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
					color.Red("No AWS profile or credentials found!")
					fmt.Println("")
					fmt.Println("Possibilities:")
					fmt.Println("- Use --profile flag or AWS_PROFILE environment variable to define an existing AWS profile")
					fmt.Println("- Use --access_key and --secret_key flag or AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variable")
					fmt.Println("")
					os.Exit(1)
				}
			}
		} else {
			os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", secretKey)
		}
	} else {
		os.Setenv("AWS_PROFILE", profile)
	}
}

func loadAwsRegion() {
	if os.Getenv("AWS_DEFAULT_REGION") != "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
}
