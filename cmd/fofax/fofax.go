package main

import (
	"fmt"
	"github.com/xiecat/fofax/internal/cli"
	"github.com/xiecat/fofax/internal/mcpserver"
	"github.com/xiecat/fofax/internal/printer"
	"github.com/xiecat/fofax/internal/runner"
)

func main() {
	option := cli.ParseOptions()
	if option.McpSSE {
		if err := mcpserver.Start(option); err != nil {
			printer.Fatal(err)
		}
		return
	}
	fofax, err := runner.NewRunner(option)
	if err != nil {
		printer.Fatal(err)
	}
	res := fofax.Run()
	res.Range(func(key, value interface{}) bool {
		if value == nil {
			fmt.Println(key)
		} else {
			fmt.Println(key, value)
		}
		return true
	})
}
