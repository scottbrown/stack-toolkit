package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

type Stack struct {
  Name string
  Region string
}

func (stack Stack) GetInstanceNames() ([]string, error) {
  config := &aws.Config{
    Region: aws.String(stack.Region),
  }

  params := &ec2.DescribeInstancesInput{
    Filters: []*ec2.Filter{
      {
        Name: aws.String("tag:aws:cloudformation:stack-name"),
        Values: []*string{
          aws.String(stack.Name),
        },
      },
    },
  }

  svc := ec2.New(session.New(), config)
  response, err := svc.DescribeInstances(params)

  if err != nil {
    return []string{}, err
  }

  instance_names := []string{}
  for _, reservation := range response.Reservations {
    for _, instance := range reservation.Instances {
      instance_names = append(instance_names, *instance.PublicDnsName)
    }
  }

  return instance_names, nil
}

