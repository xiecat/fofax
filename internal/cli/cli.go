package cli

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	FoFaXVersion = "unknown"
	Commit       = "unknown"
	Date         = "unknown"
	Branch       = "unknown"
)

func banner() {
	magenta := color.New(color.FgHiWhite).PrintfFunc()
	fmt.Println("")
	magenta(`      ____        ____       _  __` + "\n")
	magenta(`     / __/____   / __/____ _| |/ /` + "\n")
	magenta("    / /_ / __ \\ / /_ / __ `/|   / " + "\n")
	magenta(`   / __// /_/ // __// /_/ //   |  ` + "\n")
	magenta(`  /_/   \____//_/   \__,_//_/|_|  ` + "\n")
	magenta("                                    " + FoFaXVersion + "\n")
	fmt.Println("")
	g := color.New(color.FgHiGreen).PrintfFunc()
	g("                         xiecat.fun\n")
	fmt.Println("")
}
