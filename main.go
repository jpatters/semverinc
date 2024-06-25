package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
)

func main() {

	incMajor := flag.Bool("M", false, "Increment major version")
	incMinor := flag.Bool("m", false, "Increment minor version")
	incPatch := flag.Bool("p", false, "Increment patch version")

	flag.Parse()
	version := flag.Args()[0]

	v, err := semver.StrictNewVersion(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing version: %s\n", err)
		os.Exit(1)
	}

	currentVersion := *v

	if *incMajor {
		currentVersion = currentVersion.IncMajor()
	} else if *incMinor {
		currentVersion = currentVersion.IncMinor()
	} else if *incPatch {
		currentVersion = currentVersion.IncPatch()
	}
	fmt.Println(currentVersion.String())
}
