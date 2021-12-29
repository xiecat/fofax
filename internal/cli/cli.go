package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	FoFaXVersion = "unknown"
	Commit       = "unknown"
	Date         = "unknown"
	Branch       = "unknown"
)

func getVname(f string) string {
	if strings.HasPrefix(f, "v0.1.") {
		return fmt.Sprintf("Capricornus")
	} else {
		return f
	}
}

func banner() {
	magenta("")
	magenta(`      ____        ____       _  __`)
	magenta(`     / __/____   / __/____ _| |/ /`)
	magenta("    / /_ / __ \\ / /_ / __ `/|   / ")
	magenta(`   / __// /_/ // __// /_/ //   |  `)
	magenta(`  /_/   \____//_/   \__,_//_/|_|  `)
	magenta("                 	" + FoFaXVersion)
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
