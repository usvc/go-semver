package semver

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SemverTests struct {
	suite.Suite
}

func TestSemver(t *testing.T) {
	suite.Run(t, &SemverTests{})
}

func (s *SemverTests) TestSemverBumpMajor() {
	v := Semver{
		VersionCore: VersionCore{
			Major: 1,
			Minor: 2,
			Patch: 3,
		},
		PreRelease:    []string{"pr"},
		BuildMetadata: []string{"bm"},
	}
	v.BumpMajor()
	s.Equal(uint(2), v.VersionCore.Major)
	s.Equal(uint(0), v.VersionCore.Minor)
	s.Equal(uint(0), v.VersionCore.Patch)
	s.Equal(PreRelease{}, v.PreRelease)
	s.Equal(BuildMetadata{}, v.BuildMetadata)
}

func (s *SemverTests) TestSemverBumpMinor() {
	v := Semver{
		VersionCore: VersionCore{
			Major: 1,
			Minor: 2,
			Patch: 3,
		},
		PreRelease:    []string{"pr"},
		BuildMetadata: []string{"bm"},
	}
	v.BumpMinor()
	s.Equal(uint(1), v.VersionCore.Major)
	s.Equal(uint(3), v.VersionCore.Minor)
	s.Equal(uint(0), v.VersionCore.Patch)
	s.Equal(PreRelease{}, v.PreRelease)
	s.Equal(BuildMetadata{}, v.BuildMetadata)
}

func (s *SemverTests) TestSemverBumpPatch() {
	v := Semver{
		VersionCore: VersionCore{
			Major: 1,
			Minor: 2,
			Patch: 3,
		},
		PreRelease:    []string{"pr"},
		BuildMetadata: []string{"bm"},
	}
	v.BumpPatch()
	s.Equal(uint(1), v.VersionCore.Major)
	s.Equal(uint(2), v.VersionCore.Minor)
	s.Equal(uint(4), v.VersionCore.Patch)
	s.Equal(PreRelease{}, v.PreRelease)
	s.Equal(BuildMetadata{}, v.BuildMetadata)
}

func (s *SemverTests) TestSemverBumpPreRelease() {
	v := Semver{
		PreRelease: []string{
			"pr",
			"1",
		},
		BuildMetadata: []string{"bm"},
	}
	s.Nil(v.BumpPreRelease())
	s.Equal("0.0.0-pr.2", v.String())
	s.Equal(BuildMetadata{}, v.BuildMetadata)
}

func (s *SemverTests) TestSemverBumpPreRelease_lastNotNumber() {
	v := Semver{
		PreRelease: []string{
			"pr",
		},
	}
	s.NotNil(v.BumpPreRelease())
	s.Equal("0.0.0-pr", v.String())
}

func (s *SemverTests) TestSemverSetPreRelease() {
	v := Semver{}
	v.SetPreRelease("a", "b")
	s.Equal("0.0.0-a.b", v.String())
}

func (s *SemverTests) TestSemverSetBuildMetadata() {
	v := Semver{}
	v.SetBuildMetadata("a", "b")
	s.Equal("0.0.0+a.b", v.String())
}

func (s *SemverTests) TestString() {
	timestamp := time.Now().Format(time.RFC3339)
	semver := Semver{"v", VersionCore{1, 2, 3}, PreRelease{"alpha", "4"}, BuildMetadata{timestamp}}
	s.Equal(fmt.Sprintf("v1.2.3-alpha.4+%s", timestamp), semver.String())
}
