package main

import (
  "errors"
  "fmt"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

type Arguments struct {
  ProgramName string
  Region string
  StackName string
}

func (a *Arguments) Parse(arg_data []string) error {
  a.ProgramName = arg_data[0]

  if len(arg_data) < 3 {
    msg := fmt.Sprintf("Usage: %s STACK_NAME REGION", a.ProgramName)
    return errors.New(msg)
  }

  a.StackName = arg_data[1]
  a.Region = arg_data[2]

  return nil
}

func main() {
  args := Arguments{}

  err := args.Parse(os.Args)

  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }

  config := &aws.Config{
    Region: aws.String(args.Region),
  }

  svc := ec2.New(session.New(), config)

  params := &ec2.DescribeInstancesInput{
    Filters: []*ec2.Filter{
      {
        Name: aws.String("tag:aws:cloudformation:stack-name"),
        Values: []*string{
          aws.String(args.StackName),
        },
      },
    },
  }

  response, err := svc.DescribeInstances(params)

  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }

  for _, reservation := range response.Reservations {
    for _, instance := range reservation.Instances {
      fmt.Println(*instance.PublicDnsName)
    }
  }
}

