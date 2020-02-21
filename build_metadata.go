package semver

import "strings"

// BuildMetadata stores the `build-metadata` section of a
// semantic version that's appended after a plus sign ('+')
type BuildMetadata []string

// String returns a string representation of the build
// metadata
func (bm BuildMetadata) String() string {
	return strings.Join(bm, ".")
}
