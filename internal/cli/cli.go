package cli

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	FoFaXVersion = "0.1.14"
	Commit       = "unknown"
	Date         = "2021-12-27T10:26:05Z"
	Branch       = "unknown"
)

func banner() {
	magenta("")
	magenta(`      ____        ____       _  __`)
	magenta(`     / __/____   / __/____ _| |/ /`)
	magenta("    / /_ / __ \\ / /_ / __ `/|   / ")
	magenta(`   / __// /_/ // __// /_/ //   |  `)
	magenta(`  /_/   \____//_/   \__,_//_/|_|  `)
	magenta("                 	v" + FoFaXVersion)
	bannerSite("			https://fofax.xiecat.fun\n")
}
func bannerSite(a ...interface{}) {
	cl := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Fprintln(color.Error, cl(fmt.Sprint(a...)))
}
func magenta(a ...interface{}) {
	cl := color.New(color.FgHiWhite).SprintfFunc()
	fmt.Fprintln(color.Error, cl(fmt.Sprint(a...)))
}
