package semver

import (
	"regexp"
	"strconv"
	"strings"
)

// Semver stores information about a semantic version as
// defined at https://semver.org
type Semver struct {
	Prefix        string
	VersionCore   VersionCore
	PreRelease    PreRelease
	BuildMetadata BuildMetadata
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
