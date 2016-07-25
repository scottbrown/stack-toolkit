package main

import (
  "errors"
  "fmt"
  "os"
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

  region := Region{
    Name: args.Region,
  }

  stack_names, err := region.GetCreatedStacks()

  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }

  for _, name := range stack_names {
    fmt.Println(name)
  }
}

