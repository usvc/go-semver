package semver

import "strings"

type BuildMetadata []string

func (bm BuildMetadata) String() string {
	return strings.Join(bm, ".")
}
