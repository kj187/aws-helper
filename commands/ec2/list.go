package ec2

import (
	"fmt"
	"reflect"
	"github.com/apcera/termtables"
	"github.com/kj187/aws-helper/aws/ec2"
	"strings"
)

func ListInstances(region string, tags []string, filters []string, columns []string, removeColumns []string) {
	fmt.Printf("Used region: %v", region)

	if tags != nil || filters != nil {
		displayFilters := tags
		for index, _ := range filters { displayFilters = append(displayFilters, filters[index]) }
		fmt.Printf("  |  Active filter: %v\n", strings.Join(displayFilters, ", "))
	}
	if columns != nil { fmt.Printf("  |  Used additional columns: %v", columns) }
	if removeColumns != nil { fmt.Printf("  |  Removed default columns: %v", removeColumns) }

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
	instances := ec2.FindEc2InstancesByTags(region, tags, filters)

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
				if (*tag.Key == columnKey) {
					row.AddCell(*tag.Value)
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println(table.Render())	
}
