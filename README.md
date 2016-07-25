# stack-toolkit

A collection of CLI tools for use with AWS stack deployed services.

## Download

Download the latest binaries from the [Releases](https://github.com/unbounce/stack-toolkit/releases) page on this repository.  Choose the correct distribution that suits the platform where the commands will be run.

## Installation

Unzip the tar.gz file onto your local system, somewhere that the `PATH` environment variable can reach.  If unsure, run `which stacks` and it should return the path to the command.  If nothing is returned, then the files are not installed where `PATH` can see them.

## Commands

`stacks REGION`

This command lists the active CloudFormation stacks in a given `REGION`.

`stack-instances STACK_NAME REGION`

This command lists the EC2 instances that belong to a given `STACK_NAME` within a given `REGION`.

## Fun Stuff

The output from the commands can be combined with other Unix/Linux utilities to form a workflow.

### Example 1: Finding a specific group of stacks

Provide an alphabetically sorted list of stacks labeled "production" from the us-east-1 AWS region.

```
$ stacks us-east-1 | grep production | sort
```

### Example 2: SSHing into an instance

Provide the `ssh` command with one EC2 public DNS name, retrieved from the example-stack in us-east-1.  The `head -1` ensures that only one result is returned (in the case of instances behind an autoscaling group).

```
$ ssh $(stack-instances example-stack us-east-1 | head -1)
```

