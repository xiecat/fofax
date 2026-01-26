package cli

import (
	"fmt"
	"strings"

	"github.com/xiecat/fofax/internal/printer"

	"github.com/fatih/color"
)

var (
	FoFaXVersion = "0.1.48"
	Commit       = "unknown"
	Date         = "2026-01-26T15:06:09Z"
	Branch       = "unknown"
)

func getVname(f string) string {
	if strings.HasPrefix(f, "v0.1.") || strings.HasPrefix(f, "0.1.") {
		return "Capricornus"
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
	magenta("                 	v" + FoFaXVersion)
	bannerSite("			https://fofax.xiecat.fun\n")
}
func bannerSite(a ...interface{}) {
	cl := color.New(color.FgHiGreen).SprintfFunc()
	if !printer.Silent {
		fmt.Fprintln(color.Error, cl(fmt.Sprint(a...)))
	}
}
func magenta(a ...interface{}) {
	cl := color.New(color.FgHiWhite).SprintfFunc()
	if !printer.Silent {
		fmt.Fprintln(color.Error, cl(fmt.Sprint(a...)))
	}
}
