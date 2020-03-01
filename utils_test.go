package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) TestIsValid_valid() {
	validSemver := []string{
		"1.0.0",
		"v1.0.0",
		"1.0.0-pr",
		"1.0.0-pr.1",
		"1.0.0-pr.1+build",
		"1.0.0-pr.1+build.1",
	}
	for _, testSemver := range validSemver {
		s.True(IsValid(testSemver))
	}
}

func (s *UtilsTests) TestIsValid_invalid() {
	invalidSemver := []string{
		"1.0.a",
		"1.a.0",
		"a.0.0",
		"a1.0.0",
		"1.0.0-pr+k+",
		"1.0.0-pr!",
		"1.0.0-pr+",
		"1.0.0-pr+a+b",
	}
	for _, testSemver := range invalidSemver {
		s.False(IsValid(testSemver), fmt.Sprintf("%s should be invalid", testSemver))
	}
}

func (s *UtilsTests) TestParse_basic() {
	semver := Parse("1.2.3")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
}

func (s *UtilsTests) TestParse_prefixedV() {
	semver := Parse("v1.2.3")
	s.Equal("v", semver.Prefix)
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)

	semver = Parse("V1.2.3")
	s.Equal("V", semver.Prefix)
}

func (s *UtilsTests) TestParse_withPreRelease() {
	semver := Parse("1.2.3-prerel")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("prerel", semver.PreRelease.String())
}

func (s *UtilsTests) TestParse_withMultiPreRelease() {
	semver := Parse("1.2.3-prerel0.prerel1.prerel2")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("prerel0.prerel1.prerel2", semver.PreRelease.String())
}

func (s *UtilsTests) TestParse_withBuildMetadata() {
	semver := Parse("1.2.3+buildmeta")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("buildmeta", semver.BuildMetadata.String())
}

func (s *UtilsTests) TestParse_withMultiBuildMetadata() {
	semver := Parse("1.2.3+buildmeta0.buildmeta1.buildmeta2")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.EqualValues([]string{"buildmeta0", "buildmeta1", "buildmeta2"}, semver.BuildMetadata)
}

func (s *UtilsTests) TestParse_full() {
	semver := Parse("v1.2.3-prerel0.prerel1+build0.build1")
	s.Equal("v", semver.Prefix)
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.EqualValues([]string{"prerel0", "prerel1"}, semver.PreRelease)
	s.EqualValues([]string{"build0", "build1"}, semver.BuildMetadata)
}
