package bump

import (
	"flag"
	"fmt"
)

var patch = flag.Bool("patch", false, "")
var minor = flag.Bool("minor", false, "")
var major = flag.Bool("major", false, "")

func ParseFlagsToBumpType() (BumpType, error) {
	if *patch {
		return Patch, nil
	} else if *minor {
		return Minor, nil
	} else if *major {
		return Major, nil
	} else {
		return Patch, fmt.Errorf("You must specify one of --patch, --minor, --major")
	}
}
