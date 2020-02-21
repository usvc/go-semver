package semver

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PreReleaseTests struct {
	suite.Suite
}

func TestPreRelease(t *testing.T) {
	suite.Run(t, &PreReleaseTests{})
}

func (s *PreReleaseTests) assertEqual(
	expectedOrder []string,
	actualOrder PreReleases,
	messageAndArgs ...interface{},
) {
	preReleaseStrings := []string{}
	for i := 0; i < len(actualOrder); i++ {
		preReleaseStrings = append(
			preReleaseStrings,
			actualOrder[i].String(),
		)
	}
	s.Equal(
		strings.Join(expectedOrder, " , "),
		strings.Join(preReleaseStrings, " , "),
		messageAndArgs...,
	)
}

func (s *PreReleaseTests) TestString() {
	preRelease := PreRelease{"a", "b", "c"}
	s.Equal("a.b.c", preRelease.String())
}

func (s *PreReleaseTests) TestSorting_separatorDepth() {
	var prs PreReleases = []PreRelease{
		[]string{"a"},           // X.Y.Z-a
		[]string{"a", "a"},      // X.Y.Z-a.a
		[]string{"a", "a", "a"}, // X.Y.Z-a.a.a
	}
	sort.Sort(prs)
	s.assertEqual([]string{"a.a.a", "a.a", "a"}, prs)
}

func (s *PreReleaseTests) TestSorting_version() {
	var prs PreReleases = []PreRelease{
		[]string{"a", "2"},  // X.Y.Z-a.2
		[]string{"a", "1"},  // X.Y.Z-a.1
		[]string{"a", "10"}, // X.Y.Z-a.1
		[]string{"a", "3"},  // X.Y.Z-a.3
	}
	sort.Sort(prs)
	s.assertEqual([]string{"a.1", "a.2", "a.3", "a.10"}, prs)
}

func (s *PreReleaseTests) TestSorting_lexical() {
	var prs PreReleases = []PreRelease{
		[]string{"c"}, // X.Y.Z-c
		[]string{"a"}, // X.Y.Z-a
		[]string{"b"}, // X.Y.Z-b
	}
	sort.Sort(prs)
	s.assertEqual([]string{"a", "b", "c"}, prs)
}

func (s *PreReleaseTests) TestSorting_standardPreReleaseLabels() {
	var prs PreReleases = []PreRelease{
		[]string{"rc"},    // X.Y.Z-c
		[]string{"beta"},  // X.Y.Z-b
		[]string{"alpha"}, // X.Y.Z-a
	}
	sort.Sort(prs)
	s.assertEqual([]string{"alpha", "beta", "rc"}, prs)
}

func (s *PreReleaseTests) TestSorting_standardPreReleaseReleases() {
	var prs PreReleases = []PreRelease{
		[]string{"alpha", "2"},  // X.Y.Z-alpha.2
		[]string{"alpha", "10"}, // X.Y.Z-alpha.10
		[]string{"alpha"},       // X.Y.Z-alpha
		[]string{"alpha", "1"},  // X.Y.Z-alpha.1
	}
	sort.Sort(prs)
	s.assertEqual([]string{"alpha.1", "alpha.2", "alpha.10", "alpha"}, prs)
}
