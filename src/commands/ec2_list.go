package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/apcera/termtables"
	"github.com/kj187/aws-helper/src/aws/ec2"
	"github.com/spf13/cobra"
)

var tags []string
var filters []string
var columns []string
var removeColumns []string

var ec2ListCommand = &cobra.Command{
	Use:   "ec2:list",
	Short: "List all EC2 instances",
	PreRun: func(cmd *cobra.Command, args []string) {
		loadAwsRegion()
		loadAwsCredentials()
	},
	Run: func(cmd *cobra.Command, args []string) {
		listInstances(region, tags, filters, columns, removeColumns)
	},
}

func init() {
	ec2ListCommand.Flags().StringSliceVarP(&tags, "tag", "t", nil, "filter with tag (Example: Name:Jenkins)")
	ec2ListCommand.Flags().StringSliceVarP(&filters, "filter", "f", nil, "filter with column (Example: InstanceType:t2.micro)")
	ec2ListCommand.Flags().StringSliceVarP(&columns, "column", "c", nil, "add additional column (tag)")
	ec2ListCommand.Flags().StringSliceVarP(&removeColumns, "remove-column", "C", nil, "remove default column")
	RootCommand.AddCommand(ec2ListCommand)
}

func listInstances(region string, tags []string, filters []string, columns []string, removeColumns []string) {
	fmt.Printf("Used region: %v", region)

	if tags != nil || filters != nil {
		displayFilters := tags
		for index := range filters {
			displayFilters = append(displayFilters, filters[index])
		}
		fmt.Printf("  |  Active filter: %v\n", strings.Join(displayFilters, ", "))
	}
	if columns != nil {
		fmt.Printf("  |  Used additional columns: %v", columns)
	}
	if removeColumns != nil {
		fmt.Printf("  |  Removed default columns: %v", removeColumns)
	}

	defaultColumns := []string{
		"InstanceId",
		"ImageId",
		"State",
		"SubnetId",
		"AZ",
		"InstanceType",
		"KeyName",
		"PrivateIpAddress",
		"PublicIpAddress",
	}

	table := termtables.CreateTable()
	instances := ec2.GetInstances(region, tags, filters)

	for _, removeColumn := range removeColumns {
		for idx, columnName := range defaultColumns {
			if removeColumn == columnName {
				defaultColumns = append(defaultColumns[:idx], defaultColumns[idx+1:]...)
			}
		}
	}

	for _, defaultColumn := range defaultColumns {
		table.AddHeaders(defaultColumn)
	}

	for _, key := range columns {
		table.AddHeaders(key)
	}

	for _, instance := range instances {
		row := table.AddRow()

		for _, defaultColumn := range defaultColumns {
			s := reflect.ValueOf(&instance).Elem()
			row.AddCell(s.Elem().FieldByName(defaultColumn).Interface())
		}

		for _, columnKey := range columns {
			for _, tag := range instance.Tags {
				if *tag.Key == columnKey {
					row.AddCell(*tag.Value)
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println(table.Render())
}
