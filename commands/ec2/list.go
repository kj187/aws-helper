package ec2

import (
	"fmt"
	"github.com/apcera/termtables"
	"github.com/kj187/aws-helper/aws/ec2"
)

func ListInstances(region string, tags string) {
	fmt.Printf("Used region: %v\n", region)
	table := termtables.CreateTable()
	table.AddHeaders("Instance Id", "Instance Type", "KeyName", "PrivateIpAddress", "PublicIpAddress")

	instances := ec2.FindEc2InstancesByTags(region, tags)
	for _, instance := range instances {
		table.AddRow(instance.InstanceId, instance.InstanceType, instance.KeyName, instance.PrivateIpAddress, instance.PublicIpAddress)
	}

	fmt.Println("")
	fmt.Println(table.Render())	
}
