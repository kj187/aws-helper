# AWS Helper

[![AwsHelper](https://godoc.org/github.com/kj187/aws-helper?status.svg)](https://godoc.org/github.com/kj187/aws-helper)
[![Build Status](https://travis-ci.org/kj187/aws-helper.svg?branch=master)](https://travis-ci.org/kj187/aws-helper)
[![Go Report Card](https://goreportcard.com/badge/github.com/kj187/aws-helper)](https://goreportcard.com/report/github.com/kj187/aws-helper)
[![Coverage Status](https://coveralls.io/repos/github/kj187/aws-helper/badge.svg?branch=master)](https://coveralls.io/github/kj187/aws-helper?branch=master)

The AWS Helper is a go based command line interface utility for AWS

## Installation

### Linux AMD64

``` shell
$ wget -O /usr/local/bin/aws-helper XXX/artifacts/raw/aws-helper_linux-amd64
$ chmod +x /usr/local/bin/aws-helper
```

### Darwin

``` shell
$ wget -O /usr/local/bin/aws-helper XXX/artifacts/raw/aws-helper_darwin-arch386
$ chmod +x /usr/local/bin/aws-helper
```

### Contribution

In order to work properly, AWS Helper needs to be checked out at the following location: `$GOPATH/src/github.com/kj187/aws-helper`

``` shell
$ git clone XXX/aws-helper.git $GOPATH/src/github.com/kj187/aws-helper
```

#### Setup

``` shell
$ cd $GOPATH/src/github.com/kj187/aws-helper
$ make setup
```

#### Tests only

``` shell
$ cd $GOPATH/src/github.com/kj187/aws-helper
$ make test
```

## Commands

``` shell
____ _ _ _ ____    _  _ ____ _    ___  ____ ____
|__| | | | [__     |__| |___ |    |__] |___ |__/
|  | |_|_| ___]    |  | |___ |___ |    |___ |  \

The AWS Helper is a go based command line interface utility for AWS.
Author: Julian Kleinhans <mail@kj187.de>, alias @kj187

Usage:
  aws-helper [flags]
  aws-helper [command]

Examples:
aws-helper ec2:list -c Name -C KeyName -f AZ:eu-central-1

Available Commands:
  ec2:list    List all EC2 instances
  help        Help about any command

Flags:
  -a, --access_key string   set aws access_key
  -h, --help                help for aws-helper
  -p, --profile string      set aws profile
  -r, --region string       set region (default "eu-central-1")
  -s, --secret_key string   set aws secret_key

Use "aws-helper [command] --help" for more information about a command.
```

### Region

The default region is `eu-central-1`. You can define the region with an environment variable `AWS_DEFAULT_REGION` or with the flag `--region`

Example with env var:
``` shell
$ export AWS_DEFAULT_REGION=us-east-1
$ aws-helper ec2:list
```

Example with flag:
``` shell
$ aws-helper ec2:list --region us-east-1
$ aws-helper ec2:list
```

### Credentials

#### Environment variables

TODO

#### Flags

TODO

### Ec2:list

``` shell 
Usage:
  aws-helper ec2:list [flags]

Flags:
  -c, --column stringSlice          add additional column (tag)
  -f, --filter stringSlice          filter with column (Example: InstanceType:t2.micro)
  -h, --help                        help for ec2:list
  -C, --remove-column stringSlice   remove default column
  -t, --tag stringSlice             filter with tag (Example: Name:Jenkins)

Global Flags:
  -a, --access_key string   set aws access_key
  -p, --profile string      set aws profile
  -r, --region string       set region (default "eu-central-1")
  -s, --secret_key string   set aws secret_key
```

#### Default columns

* InstanceId
* ImageId
* State
* SubnetId
* AZ
* InstanceType
* KeyName
* PrivateIpAddress
* PublicIpAddress

With the uppercase C flag `-C` or `--remove-column` you have the possibility to remove a default column.

Example 
``` shell 
$ aws-helper ec2:list -C SubnetId -C AZ
```

As you can see you could also remove multiple columns.

#### Add tags as column

Imagine your instances have a "Name" tag, with the lowercase `-c` or `--column` yu have the possibility to add tags as a column.

Example 
``` shell 
$ aws-helper ec2:list -c Name
```

Yes, you could also add multiple columns. 

#### Filter

There a two different ways to filter your results. You could use a tag filter or a column filter.

##### Tag filter

Example 
``` shell 
$ aws-helper ec2:list -t Name:Jenkins
```

##### Column filter

Example 
``` shell 
$ aws-helper ec2:list -f InstanceType:t2.micro
```
