package main

import (
	"fmt"
	"fofax/internal/cli"
	"fofax/internal/printer"
	"fofax/internal/runner"
	"os"
)

func main() {
	option := cli.ParseOptions()
	fofax, err := runner.NewRunner(option)
	if err != nil {
		printer.Error(err)
		os.Exit(1)
	}
	res :=fofax.Run()
	res.Range(func(key, value interface{}) bool {
		if value == nil {
			fmt.Println(key)
		} else {
			fmt.Println(key, value)
		}
		return true
	})
}
