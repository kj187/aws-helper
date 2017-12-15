package ec2

import (
	"fmt"
	"log"
	
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/apcera/termtables"
)

func ListInstances(region string, filter string) {
	sess := session.Must(session.NewSession())
	fmt.Printf("Used region: %v\n", region)

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeInstances(getFilter(filter))
	if err != nil {
		fmt.Println("there was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	//fmt.Printf("%+v\n", *resp)

	table := termtables.CreateTable()
	table.AddHeaders("Instance Id", "Instance Type", "KeyName", "PrivateIpAddress", "PublicIpAddress")

	for idx, _ := range resp.Reservations {
		//fmt.Println("  > Reservation Id", *res.ReservationId, " Num Instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			table.AddRow(*inst.InstanceId, *inst.InstanceType, *inst.KeyName, *inst.PrivateIpAddress, *inst.PublicIpAddress)
		}
	}

	fmt.Println("")
	fmt.Println(table.Render())
}

func getFilter(filter string) *ec2.DescribeInstancesInput {
	var describeInstancesInput *ec2.DescribeInstancesInput
	if filter != "" {
		fmt.Printf("Active filter: %v\n", filter)
		var filters []*ec2.Filter
		splittedFilterResults := strings.Split(filter, ",")
		for i := range splittedFilterResults {
			tags := strings.Split(strings.TrimSpace(splittedFilterResults[i]), ":")
			filter := &ec2.Filter {
				Name: aws.String("tag:" + tags[0]),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tags[1], "*"}, "")),
				},
			}
			filters = append(filters, filter)
		}
		describeInstancesInput = &ec2.DescribeInstancesInput{
			Filters: filters,
		}
	}	

	return describeInstancesInput
}