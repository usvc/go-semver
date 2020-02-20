package semver

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SemverTests struct {
	suite.Suite
}

func TestSemver(t *testing.T) {
	suite.Run(t, &SemverTests{})
}

func (s *SemverTests) TestParse_basic() {
	semver := Parse("1.2.3")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
}

func (s *SemverTests) TestParse_prefixedV() {
	semver := Parse("v1.2.3")
	s.Equal("v", semver.Prefix)
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)

	semver = Parse("V1.2.3")
	s.Equal("V", semver.Prefix)
}

func (s *SemverTests) TestParse_withPreRelease() {
	semver := Parse("1.2.3-prerel")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("prerel", semver.PreRelease.String())
}

func (s *SemverTests) TestParse_withMultiPreRelease() {
	semver := Parse("1.2.3-prerel0.prerel1.prerel2")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("prerel0.prerel1.prerel2", semver.PreRelease.String())
}

func (s *SemverTests) TestParse_withBuildMetadata() {
	semver := Parse("1.2.3+buildmeta")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("buildmeta", semver.BuildMetadata.String())
}

func (s *SemverTests) TestParse_withMultiBuildMetadata() {
	semver := Parse("1.2.3+buildmeta0.buildmeta1.buildmeta2")
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("buildmeta0.buildmeta1.buildmeta2", semver.BuildMetadata.String())
}

func (s *SemverTests) TestParse_full() {
	semver := Parse("v1.2.3-prerel0.prerel1+build0.build1")
	s.Equal("v", semver.Prefix)
	s.EqualValues(1, semver.VersionCore.Major)
	s.EqualValues(2, semver.VersionCore.Minor)
	s.EqualValues(3, semver.VersionCore.Patch)
	s.Equal("prerel0.prerel1", semver.PreRelease.String())
	s.Equal("build0.build1", semver.BuildMetadata.String())
}
