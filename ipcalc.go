package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"github.com/vanytsvetkov/ipcalc/functions"
)

func main() {
	var prefix, separator string
	var subnetSize int

	flag.IntVarP(&subnetSize, "split", "s", 0, "Division prefix by {int-arg} subnets.")
	flag.StringVarP(&separator, "delimiter", "d", "\n", "Join subnets by {str-arg} delimiter.")

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: ipcalc <prefix>")
		flag.PrintDefaults()
		return
	}

	prefix = flag.Arg(0)

	ipcalcResult, err := functions.IPcalc(prefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ipcalcResult)

	if subnetSize > 0 {
		ipsplitResult, err := functions.IPsplit(prefix, subnetSize, separator)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ipsplitResult)
	}

}
