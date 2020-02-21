package semver

import "fmt"

// VersionCore represents the major, minor, and patch
// versions section of a semantic version
type VersionCore struct {
	Major uint
	Minor uint
	Patch uint
}

// String returns the string representation of the VersionCore structure
func (vc VersionCore) String() string {
	return fmt.Sprintf("%v.%v.%v", vc.Major, vc.Minor, vc.Patch)
}

// VersionCores is a sortable slice of VersionCore structures
type VersionCores []VersionCore

// Len implements sort.Interface's Len()
func (vcs VersionCores) Len() int {
	return len(vcs)
}

// Swap implements sort.Interface's Swap()
func (vcs VersionCores) Swap(i, j int) {
	vcs[i], vcs[j] = vcs[j], vcs[i]
}

// Less implements sort.Interface's Less()
func (vcs VersionCores) Less(i, j int) bool {
	switch true {
	case vcs[i].Major < vcs[j].Major:
		fallthrough
	case vcs[i].Major == vcs[j].Major && vcs[i].Minor < vcs[j].Minor:
		fallthrough
	case vcs[i].Major == vcs[j].Major && vcs[i].Minor == vcs[j].Minor && vcs[i].Patch < vcs[j].Patch:
		return true
	default:
		return false
	}
}
