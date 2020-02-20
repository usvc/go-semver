package semver

type Semvers []Semver

func (semvers Semvers) Len() int {
	return len(semvers)
}

func (semvers Semvers) Swap(i, j int) {
	semvers[i], semvers[j] = semvers[j], semvers[i]
}

func (semvers Semvers) Less(i, j int) bool {
	vcs := VersionCores{semvers[i].VersionCore, semvers[j].VersionCore}
	versionCoreIsLess := vcs.Less(0, 1)
	return versionCoreIsLess
}
