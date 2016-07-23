package main

import (
  "errors"
  "fmt"
  "os"
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

  a.Region = arg_data[1]
  a.StackName = arg_data[2]

  return nil
}

func main() {
  args := Arguments{}

  err := args.Parse(os.Args)

  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }
}

