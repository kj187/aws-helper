package ec2

import (
	"strings"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"	
)

type Instance struct {
	_ struct{} `type:"structure"`

	InstanceId string `type:"string"`
	InstanceType string `type:"string"`
	KeyName string `type:"string"`
	PrivateIpAddress string `type:"string"`
	PublicIpAddress string `type:"string"`
}

func FindEc2InstancesByTags(region string, tags string) []*Instance {
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeInstances(getTags(tags))
	if err != nil {
		fmt.Println("there was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	//fmt.Printf("%+v\n", *resp)

	var instances []*Instance 
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			instance := Instance{
				InstanceId: *inst.InstanceId,
				InstanceType: *inst.InstanceType,
				KeyName: *inst.KeyName,
				PrivateIpAddress: *inst.PrivateIpAddress,
				PublicIpAddress: *inst.PublicIpAddress,
			 }

			 instances = append(instances, &instance)
		}
	}

	return instances
}

func getTags(tags string) *ec2.DescribeInstancesInput {
	var describeInstancesInput *ec2.DescribeInstancesInput
	if tags != "" {
		fmt.Printf("Active filter: %v\n", tags)
		var filters []*ec2.Filter
		splittedTags := strings.Split(tags, ",")
		for i := range splittedTags {
			tag := strings.Split(strings.TrimSpace(splittedTags[i]), ":")
			filter := &ec2.Filter {
				Name: aws.String("tag:" + tag[0]),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tag[1], "*"}, "")),
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