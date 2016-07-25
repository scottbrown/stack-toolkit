package main

import (
  "errors"
  "fmt"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

type Arguments struct {
  ProgramName string
  Region string
}

func (a *Arguments) Parse(arg_data []string) error {
  a.ProgramName = arg_data[0]

  if len(arg_data) < 2 {
    msg := fmt.Sprintf("Usage: %s REGION", a.ProgramName)
    return errors.New(msg)
  }

  a.Region = arg_data[1]

  return nil
}

func main() {
  args := Arguments{}

  err := args.Parse(os.Args)

  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }

  svc := cloudformation.New(session.New(), &aws.Config{Region: aws.String(args.Region)})

  svc_input := cloudformation.DescribeStacksInput{}
  response, err := svc.DescribeStacks(&svc_input)

  if err != nil {
    fmt.Fprintln(os.Stderr, "%s\n", err.Error())
  }

  for i := range response.Stacks {
    fmt.Println(i)
  }
}

