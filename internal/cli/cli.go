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
	magenta := color.New(color.FgMagenta).PrintfFunc()
	fmt.Println("")
	magenta("    ____      ____          \n")
	magenta(`   / __/___  / __/___ __  __` + "\n")
	magenta(`  / /_/ __ \/ /_/ __ ` + "`/ |/_/" + "\n")
	magenta(` / __/ /_/ / __/ /_/ />  <  ` + "\n")
	magenta(`/_/  \____/_/  \__,_/_/|_|  ` + "\n")
	fmt.Println("")
	g := color.New(color.FgGreen).PrintfFunc()
	g(fmt.Sprintf("                                   { %s  Author:xiecat }\n", FoFaXVersion))
	fmt.Println("")
}
