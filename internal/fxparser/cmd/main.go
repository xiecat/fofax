package main

import (
	"fmt"

	"github.com/xiecat/fofax/internal/fxparser"
)

func main() {
	query := "host=\".gov.cn\"&&fx=\"aa\"&&(ip=\"1.1.1.1中国\")&&(after=\"2017\" && before=\"2017-10-01\")"
	//query = `(host=".gov.cn"&&fx="aa")&&(ip="1.1.1.1")&&(after="2017" && before="2017-10-01")`
	fmt.Println(fxparser.Query(query))
}
