package cli

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	FoFaXVersion = "v1.0"
	Commit       = ""
	Date         = ""
	Branch       = ""
)

func banner() {
	magenta := color.New(color.FgMagenta).PrintfFunc()
	fmt.Println("")
	magenta("      _____     _____       _____     _____      __  __   \n")
	magenta(`    /\_____\   ) ___ (    /\_____\   /\___/\   /\  /\  /\ ` + "\n")
	magenta(`   ( (  ___/  / /\_/\ \  ( (  ___/  / / _ \ \  \ \ \/ / / ` + "\n")
	magenta(`    \ \ \_   / /_/ (_\ \  \ \ \_    \ \(_)/ /   \ \  / /  ` + "\n")
	magenta(`    / / /_\  \ \ )_/ / /  / / /_\   / / _ \ \   / /  \ \  ` + "\n")
	magenta(`   / /____/   \ \/_\/ /  / /____/  ( (_( )_) ) / / /\ \ \ ` + "\n")
	magenta(`   \/_/        )_____(   \/_/       \/_/ \_\/  \/__\/__\/ ` + "\n")
	fmt.Println("")
	g := color.New(color.FgGreen).PrintfFunc()
	g(fmt.Sprintf("                                   { %s  Author:xiecat }\n", FoFaXVersion))
	fmt.Println("")

}
