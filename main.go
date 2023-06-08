package main

import (
	"flag"
	"fmt"

	"github.com/benjlevesque/gh-bump/pkg/bump"
	"github.com/benjlevesque/gh-bump/pkg/gh"
)

var dryRun = flag.Bool("dryRun", false, "")
var title = flag.String("title", "", "the release title")

func main() {
	flag.Parse()

	bumpType, err := bump.ParseFlagsToBumpType()
	if err != nil {
		fmt.Println(err)
		return
	}

	latest, err := gh.GetLatestRelease()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("latest is %s, ", latest)

	newTag := bump.GetNewTag(latest, bumpType)
	fmt.Printf("next is %s\n", newTag)

	if *dryRun {
		fmt.Printf("[Dry Run] creating release for %s\n", newTag)
		return
	}
	release, err := gh.CreateRelease(newTag, *title)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(release)
}
