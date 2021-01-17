package main

import (
	"fmt"
	"flag"
	"os"
)

const banner = `
 _               _                                            __ _ _ _
| |__   ___  ___| |_       _____      ___ __   ___ _ __      / _(_) | |_ ___ _ __
| '_ \ / _ \/ __| __|____ / _ \ \ /\ / / '_ \ / _ \ '__|____| |_| | | __/ _ \ '__|
| | | | (_) \__ \ ||_____| (_) \ V  V /| | | |  __/ | |_____|  _| | | ||  __/ |
|_| |_|\___/|___/\__|     \___/ \_/\_/ |_| |_|\___|_|       |_| |_|_|\__\___|_|
==================================================================================
`
const Version = `0.1`




func showBanner() {
	fmt.Printf("%s", banner)
	fmt.Printf("\t\t\t\t\t\t\t\tversion: %s\n\n",Version)
}

func (options *hof.Options) validateOptions() { 

	if len(options.Hosts) > 0 {
		_, err := os.Stat(options.Hosts)
		if os.IsNotExist(err) {
			fmt.Printf("[x] Input file does not exist: %s",options.Hosts)
			os.Exit(3)
		}
	}
	if len(options.Domainf) > 0 {
		_, err := os.Stat(options.Domainf)
		if os.IsNotExist(err) {
			fmt.Printf("[x] Input file does not exist: %s",options.Domainf)
			os.Exit(3)
		}
	}

	if len(options.SoaKbFile) > 0 {
		_, err := os.Stat(options.SoaKbFile)
		if os.IsNotExist(err) {
			fmt.Printf("[x] Input file does not exist: %s",options.SoaKbFile)
			os.Exit(3)
		}
	}
	
}

func parseOptions() *hof.Options {
	gpath := os.Getenv("GOPATH")
	ProjectDir := "/src/githuib.com/dogasantos/host-owner-filter/"
	soaKbFile := "soakb.json"
	defaultConfig := gpath + ProjectDir + soaKbFile

	options := &hof.Options{}

	flag.StringVar(&options.Domainf, 		"K", "", "List of known domains to serve as a seed to compare")
	flag.StringVar(&options.Hosts, 			"L", "", "File input with list of subdomains")
	flag.StringVar(&options.OutputFile, 	"o", "", "File to write output to (optional)")
	flag.StringVar(&options.SoaKbFile, 		"s", defaultConfig, "Soa servers to filter out (a default one is provided if not set)")
	flag.BoolVar(&options.Verbose, 			"verbose", false, "Verbose output")
	
	flag.Parse()

	showBanner()
	options.validateOptions()
	return options
}

func main() {	
	options := parseOptions()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	host-ower-filter.process(options)

}