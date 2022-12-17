package main

import (
	"fmt"
	"github.com/lewinloo/cloud-station/cli"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
