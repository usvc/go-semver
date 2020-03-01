package semver

import (
	"regexp"
	"strconv"
	"strings"
)

// IsValid checks whether the provided :semver matches the
// semver grammar specification
func IsValid(semver string) bool {
	matcher := regexp.MustCompile(RegexpWithCaptureGroup)
	return matcher.MatchString(semver)
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
