package bump

import "github.com/coreos/go-semver/semver"

type BumpType int64

const (
	Patch BumpType = iota
	Minor
	Major
)

func GetNewTag(current string, bumpType BumpType) string {
	prependV := false
	if current[0] == 'v' {
		current = current[1:]
		prependV = true
	}
	ver := semver.New(current)
	switch bumpType {
	case Patch:
		ver.BumpPatch()
	case Minor:
		ver.BumpMinor()
	case Major:
		ver.BumpMajor()
	}
	newTag := ver.String()
	if prependV {
		newTag = "v" + newTag
	}

	return newTag
}
