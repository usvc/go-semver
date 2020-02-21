package semver

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type VersionCoreTests struct {
	suite.Suite
}

func TestVersionCore(t *testing.T) {
	suite.Run(t, &VersionCoreTests{})
}

func (s *VersionCoreTests) TestString() {
	versionCore := VersionCore{1, 2, 3}
	s.Equal("1.2.3", versionCore.String())
}

func (s *VersionCoreTests) TestSorting() {
	var vcs VersionCores = []VersionCore{
		{10, 10, 10},
		{10, 10, 9},
		{10, 9, 10},
		{10, 9, 9},
		{1, 9, 9},
		{1, 9, 1},
		{1, 1, 9},
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	sort.Sort(vcs)
}
