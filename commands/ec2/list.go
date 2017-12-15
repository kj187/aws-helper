package ec2

import (
	"fmt"
	"log"
	
//	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/apcera/termtables"
)

func ListInstances() {
	sess := session.Must(session.NewSession())

	nameFilter := "Dashboard" //os.Args[1]
	awsRegion := "eu-central-1"
	svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, awsRegion)

	params := &ec2.DescribeInstancesInput{
	/*
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	*/		

	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
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