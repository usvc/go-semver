package semver

import (
	"regexp"
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
}

// BumpMinor bumps the minor version
func (semver *Semver) BumpMinor() {
	semver.VersionCore.Minor++
}

// BumpPatch bumps the patch version
func (semver *Semver) BumpPatch() {
	semver.VersionCore.Patch++
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

// Parse receives a string and returns a Semver instance that
// represents the semantic version
func Parse(semver string) *Semver {
	matcher := regexp.MustCompile(RegexpWithCaptureGroup)
	matchMap := map[string]string{}
	subexpNames := matcher.SubexpNames()
	submatches := matcher.FindStringSubmatch(semver)
	for i := 1; i < len(submatches); i++ {
		matchMap[subexpNames[i]] = submatches[i]
	}
	retval := &Semver{}
	if val, ok := matchMap["Prefix"]; ok && len(val) > 0 {
		retval.Prefix = val
	}
	if val, ok := matchMap["Major"]; ok {
		if majorVersion, err := strconv.ParseUint(val, parseDecimal, parse32bit); err == nil {
			retval.VersionCore.Major = uint(majorVersion)
		}
	}
	if val, ok := matchMap["Minor"]; ok {
		if minorVersion, err := strconv.ParseUint(val, parseDecimal, parse32bit); err == nil {
			retval.VersionCore.Minor = uint(minorVersion)
		}
	}
	if val, ok := matchMap["Patch"]; ok {
		if patchVersion, err := strconv.ParseUint(val, 10, 8); err == nil {
			retval.VersionCore.Patch = uint(patchVersion)
		}
	}
	if val, ok := matchMap["PreRelease"]; ok && len(val) > 0 {
		retval.PreRelease = strings.Split(val, ".")
	}
	if val, ok := matchMap["BuildMetadata"]; ok && len(val) > 0 {
		retval.BuildMetadata = strings.Split(val, ".")
	}
	return retval
}
