package semver

import (
	"regexp"
	"strconv"
	"strings"
)

type Semver struct {
	Prefix        string
	VersionCore   VersionCore
	PreRelease    PreRelease
	BuildMetadata BuildMetadata
}

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

// Parse receives a string and returns a Semver structure that represents
// the string
func Parse(semver string) *Semver {
	// regex adapted from https://semver.org/
	matcher := regexp.MustCompile(`(?:(?P<Prefix>[vV]))?(?P<Major>0|[1-9]\d*)\.(?P<Minor>0|[1-9]\d*)\.(?P<Patch>0|[1-9]\d*)(?:-(?P<PreRelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<BuildMetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
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
		if majorVersion, err := strconv.ParseUint(val, 10, 8); err == nil {
			retval.VersionCore.Major = uint(majorVersion)
		}
	}
	if val, ok := matchMap["Minor"]; ok {
		if minorVersion, err := strconv.ParseUint(val, 10, 8); err == nil {
			retval.VersionCore.Minor = uint(minorVersion)
		}
	}
	if val, ok := matchMap["Patch"]; ok {
		if patchVersion, err := strconv.ParseUint(val, 10, 8); err == nil {
			retval.VersionCore.Patch = uint(patchVersion)
		}
	}
	if val, ok := matchMap["PreRelease"]; ok && len(val) > 0 {
		retval.PreRelease = []string{val}
	}
	if val, ok := matchMap["BuildMetadata"]; ok && len(val) > 0 {
		retval.BuildMetadata = []string{val}
	}
	return retval
}
