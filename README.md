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

