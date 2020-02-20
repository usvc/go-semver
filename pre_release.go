package semver

import "strings"

type PreRelease []string

func (pr PreRelease) String() string {
	return strings.Join(pr, ".")
}
