package semver

import "fmt"

type VersionCore struct {
	Major uint
	Minor uint
	Patch uint
}

func (vc VersionCore) String() string {
	return fmt.Sprintf("%v.%v.%v", vc.Major, vc.Minor, vc.Patch)
}

type VersionCores []VersionCore

func (vcs VersionCores) Len() int {
	return len(vcs)
}

func (vcs VersionCores) Swap(i, j int) {
	vcs[i], vcs[j] = vcs[j], vcs[i]
}

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
