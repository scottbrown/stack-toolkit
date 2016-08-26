package main

import (
  "fmt"
  "os"
  "time"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "stacks"
  app.Copyright = "(c) 2016 Unbounce Marketing Solutions Inc."
  app.Compiled = time.Now()
  app.UsageText = "stacks REGION"
  app.Version = Version
  app.Authors = []cli.Author{
    cli.Author{
      Name: "Infrastructure Team",
    },
  }
  app.Usage = "Lists all active/created stacks in a given AWS region."
  app.Action = func(c *cli.Context) error {
    name := ""
    if c.NArg() >= 1 {
      name = c.Args()[0]
    } else {
      cli.ShowAppHelp(c)
      return nil
    }

    if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
      fmt.Fprintf(os.Stderr, "%s\n", "Please set an AWS_ACCESS_KEY_ID")
      os.Exit(1)
    }

    if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
      fmt.Fprintf(os.Stderr, "%s\n", "Please set an AWS_SECRET_ACCESS_KEY")
      os.Exit(1)
    }

    region := Region{ Name: name }

    stack_names, err := region.GetCreatedStacks()

    if err != nil {
      fmt.Fprintf(os.Stderr, "%s\n", err.Error())
      os.Exit(1)
    }

    for _, name := range stack_names {
      fmt.Println(name)
    }

    return nil
  }

  app.Run(os.Args)
}

