package ec2

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	_ struct{} `type:"structure"`

	InstanceId       string     `type:"string"`
	ImageId          string     `type:"string"`
	State            string     `type:"string"`
	SubnetId         string     `type:"string"`
	AZ               string     `type:"string"`
	InstanceType     string     `type:"string"`
	KeyName          string     `type:"string"`
	PrivateIpAddress string     `type:"string"`
	PublicIpAddress  string     `type:"string"`
	Tags             []*ec2.Tag `type:"list"`
}

func GetInstances(region string, tags []string, filters []string) []*Instance {
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeInstances(buildDescribeInstancesInput(tags, filters))
	if err != nil {
		fmt.Println("there was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	//fmt.Printf("%+v\n", *resp)

	var instances []*Instance
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {

			// Workaround because PublicIpAddress is not available sometimes
			publicIpAddress := ""
			if *inst.PublicDnsName != "" {
				publicIpAddress = *inst.PublicIpAddress
			}

			instance := Instance{
				InstanceId:       *inst.InstanceId,
				ImageId:          *inst.ImageId,
				State:            *inst.State.Name,
				SubnetId:         *inst.SubnetId,
				AZ:               *inst.Placement.AvailabilityZone,
				InstanceType:     *inst.InstanceType,
				KeyName:          *inst.KeyName,
				PrivateIpAddress: *inst.PrivateIpAddress,
				PublicIpAddress:  publicIpAddress,
				Tags:             inst.Tags,
			}

			instances = append(instances, &instance)
		}
	}

	return instances
}

func buildDescribeInstancesInput(tags []string, userFilters []string) *ec2.DescribeInstancesInput {
	var filters []*ec2.Filter

	if tags != nil {
		for i := range tags {
			tag := strings.Split(strings.TrimSpace(tags[i]), ":")
			filter := &ec2.Filter{
				Name: aws.String("tag:" + tag[0]),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tag[1], "*"}, "")),
				},
			}
			filters = append(filters, filter)
		}
	}

	if userFilters != nil {
		filterMapping := make(map[string]string)
		filterMapping["instanceid"] = "instance-id"
		filterMapping["imageid"] = "image-id"
		filterMapping["state"] = "state"
		filterMapping["subnetid"] = "network-interface.subnet-id"
		filterMapping["az"] = "availability-zone"
		filterMapping["instancetype"] = "instance-type"
		filterMapping["keyname"] = "key-name"
		filterMapping["privateipaddress"] = "private-ip-address"
		filterMapping["publicipaddress"] = "public-ip-address"

		for i := range userFilters {
			x := strings.Split(strings.TrimSpace(userFilters[i]), ":")
			name := strings.ToLower(x[0])
			filter := &ec2.Filter{
				Name: aws.String(filterMapping[name]),
				Values: []*string{
					aws.String(strings.Join([]string{"*", x[1], "*"}, "")),
				},
			}
			filters = append(filters, filter)
		}
	}

	return &ec2.DescribeInstancesInput{
		Filters: filters,
	}
}
