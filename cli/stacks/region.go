package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

type Region struct {
  Name string
}

func (region Region) GetCreatedStacks() ([]string, error) {
  config := &aws.Config{
    Region: aws.String(region.Name),
  }

  svc := cloudformation.New(session.New(), config)
  response, err := svc.DescribeStacks(nil)

  if err != nil {
    return []string{}, err
  }

  stack_names := []string{}
  for _, stack := range response.Stacks {
    stack_names = append(stack_names, *stack.StackName)
  }

  return stack_names, nil
}

