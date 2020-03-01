package semver

import (
	"strconv"
	"strings"
)

// Semver stores information about a semantic version as
// defined at https://semver.org
type Semver struct {
	Prefix        string        `json:"prefix"`
	VersionCore   VersionCore   `json:"version_core"`
	PreRelease    PreRelease    `json:"pre_release"`
	BuildMetadata BuildMetadata `json:"build_metadata"`
}

// BumpMajor bumps the major version
func (semver *Semver) BumpMajor() {
	semver.VersionCore.Major++
	semver.VersionCore.Minor = 0
	semver.VersionCore.Patch = 0
	semver.PreRelease = []string{}
	semver.BuildMetadata = []string{}
}

// BumpMinor bumps the minor version
func (semver *Semver) BumpMinor() {
	semver.VersionCore.Minor++
	semver.VersionCore.Patch = 0
	semver.PreRelease = []string{}
	semver.BuildMetadata = []string{}
}

// BumpPatch bumps the patch version
func (semver *Semver) BumpPatch() {
	semver.VersionCore.Patch++
	semver.PreRelease = []string{}
	semver.BuildMetadata = []string{}
}

// BumpPreRelease bumps the pre-release version if applicable,
// if pre-release does not have a numerical version, returns
// an error
func (semver *Semver) BumpPreRelease() error {
	preReleaseVersionString := semver.PreRelease[len(semver.PreRelease)-1]
	preReleaseVersion, err := strconv.ParseUint(preReleaseVersionString, parseDecimal, parse32bit)
	if err != nil {
		return err
	}
	preReleaseVersion++
	semver.PreRelease[len(semver.PreRelease)-1] = strconv.Itoa(int(preReleaseVersion))
	semver.BuildMetadata = []string{}
	return nil
}

// SetPreRelease sets the pre-release section of the
// semantic version
func (semver *Semver) SetPreRelease(labels ...string) {
	semver.PreRelease = PreRelease(labels)
}

// SetBuildMetadata sets the build-metadata section of the
// semantic version
func (semver *Semver) SetBuildMetadata(labels ...string) {
	semver.BuildMetadata = BuildMetadata(labels)
}

// String returns the string representation of this instance of
// the Semver struct
func (semver Semver) String() string {
	builder := strings.Builder{}
	if len(semver.Prefix) > 0 {
		builder.WriteString(semver.Prefix)
	}
	builder.WriteString(semver.VersionCore.String())
	if len(semver.PreRelease) > 0 {
		builder.WriteByte('-')
		builder.WriteString(semver.PreRelease.String())
	}
	if len(semver.BuildMetadata) > 0 {
		builder.WriteByte('+')
		builder.WriteString(semver.BuildMetadata.String())
	}
	return builder.String()
}
