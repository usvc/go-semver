package semver

import (
	"math"
	"strconv"
	"strings"
)

// PreRelease defines the pre-release label often appended
// after a hypen ('-') in a semantic version
type PreRelease []string

// String returns the string representation of the pre-
// release label
func (pr PreRelease) String() string {
	return strings.Join(pr, ".")
}

// PreReleases is a sortable slice of PreRelease structures
type PreReleases []PreRelease

// Len implements sort.Interface's Len()
func (prs PreReleases) Len() int { return len(prs) }

// Swap implements sort.Interface's Swap()
func (prs PreReleases) Swap(i, j int) { prs[i], prs[j] = prs[j], prs[i] }

// Less implements sort.Interface's Less()
func (prs PreReleases) Less(i, j int) bool {
	// name these properly for future sanity
	preReleaseI := prs[i]
	preReleaseJ := prs[j]
	depthOfI := len(preReleaseI)
	depthOfJ := len(preReleaseJ)
	deepestDepth := int(math.Max(float64(depthOfI), float64(depthOfJ)))
	var iM, jM int64
	var err error
	for m := 0; m < deepestDepth; m++ {
		switch true {
		case m == depthOfI:
			return false
		case m == depthOfJ:
			return true
		case preReleaseI[m] != preReleaseJ[m]:
			if iM, err = strconv.ParseInt(
				preReleaseI[m],
				parseDecimal,
				parse32bit,
			); err != nil {
				return preReleaseI[m] < preReleaseJ[m]
			} else if jM, err = strconv.ParseInt(
				preReleaseJ[m],
				parseDecimal,
				parse32bit,
			); err != nil {
				return preReleaseI[m] < preReleaseJ[m]
			}
			return iM < jM
		}
	}
	return false
}
