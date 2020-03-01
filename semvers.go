package semver

// Semvers is a slice of Semver structures with sorting capabilities
type Semvers []*Semver

// Len implements sort.Interface's Len()
func (semvers Semvers) Len() int { return len(semvers) }

// Swap implements sort.Interface's Swap()
func (semvers Semvers) Swap(i, j int) { semvers[i], semvers[j] = semvers[j], semvers[i] }

// Less implements sort.Interface's Less()
func (semvers Semvers) Less(i, j int) bool {
	vcs := VersionCores{semvers[i].VersionCore, semvers[j].VersionCore}
	if vcs.Less(0, 1) {
		return true
	}
	prs := PreReleases{semvers[i].PreRelease, semvers[j].PreRelease}
	if prs.Less(0, 1) {
		return true
	}
	return false
}
