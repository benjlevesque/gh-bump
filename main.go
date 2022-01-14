package main

import (
	"flag"
	"fmt"

	"github.com/coreos/go-semver/semver"
)

var patch = flag.Bool("patch", false, "")
var minor = flag.Bool("minor", false, "")
var major = flag.Bool("major", false, "")

func main() {
	flag.Parse()
	latest, err := getLatest()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("latest is %s, ", latest)

	prependV := false
	if latest[0] == 'v' {
		latest = latest[1:]
		prependV = true
	}
	ver := semver.New(latest)
	if *patch {
		ver.BumpPatch()
	} else if *minor {
		ver.BumpMinor()
	} else if *major {
		ver.BumpMajor()
	} else {
		fmt.Println("You must specify one of --patch, --minor, --major")
		return
	}
	newTag := ver.String()
	if prependV {
		newTag = "v" + newTag
	}
	fmt.Printf("next is %s\n", newTag)
	release, err := newRelease(newTag)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(release)
}
