package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
)

const usage = `Usage: semver [options] <version>
    -M, --major               Increment the major version.
    -m, --minor               Increment the minor version.
    -p, --patch               Increment the patch version.`

func main() {

	var incMajor, incMinor, incPatch bool

	flag.BoolVar(&incMajor, "M", false, "Increment major version")
	flag.BoolVar(&incMajor, "major", false, "Increment major version")
	flag.BoolVar(&incMinor, "m", false, "Increment minor version")
	flag.BoolVar(&incMinor, "minor", false, "Increment minor version")
	flag.BoolVar(&incPatch, "p", false, "Increment patch version")
	flag.BoolVar(&incPatch, "patch", false, "Increment patch version")

	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if incMajor && incMinor || incMajor && incPatch || incMinor && incPatch {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Error: Only one of -M, -m, or -p can be specified\n")
		os.Exit(1)
	}

	if !incMajor && !incMinor && !incPatch {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Error: One of -M, -m, or -p must be specified\n")
		os.Exit(1)
	}

	version := flag.Args()[0]

	v, err := semver.StrictNewVersion(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing version: %s\n", err)
		os.Exit(1)
	}

	currentVersion := *v

	if incMajor {
		currentVersion = currentVersion.IncMajor()
	} else if incMinor {
		currentVersion = currentVersion.IncMinor()
	} else if incPatch {
		currentVersion = currentVersion.IncPatch()
	}
	fmt.Println(currentVersion.String())
}
