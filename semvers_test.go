package semver

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SemversTests struct {
	suite.Suite
}

func TestSemvers(t *testing.T) {
	suite.Run(t, &SemversTests{})
}

func (s *SemversTests) assertEqual(
	expectedOrder []string,
	actualOrder Semvers,
	messageAndArgs ...interface{},
) {
	semverStrings := []string{}
	for i := 0; i < len(actualOrder); i++ {
		semverStrings = append(
			semverStrings,
			actualOrder[i].String(),
		)
	}
	s.Equal(
		strings.Join(expectedOrder, " , "),
		strings.Join(semverStrings, " , "),
		messageAndArgs...,
	)
}

func (s *SemversTests) TestSorting_majorVersion() {
	semvers := Semvers{
		Parse("10.0.0"),
		Parse("3.0.0"),
		Parse("2.0.0"),
		Parse("1.0.0"),
	}
	sort.Sort(semvers)
	s.assertEqual([]string{"1.0.0", "2.0.0", "3.0.0", "10.0.0"}, semvers)
}

func (s *SemversTests) TestSorting_minorVersion() {
	semvers := Semvers{
		Parse("0.10.0"),
		Parse("0.3.0"),
		Parse("0.2.0"),
		Parse("0.1.0"),
	}
	sort.Sort(semvers)
	s.assertEqual([]string{"0.1.0", "0.2.0", "0.3.0", "0.10.0"}, semvers)
}

func (s *SemversTests) TestSorting_patchVersion() {
	semvers := Semvers{
		Parse("0.0.10"),
		Parse("0.0.3"),
		Parse("0.0.2"),
		Parse("0.0.1"),
	}
	sort.Sort(semvers)
	s.assertEqual([]string{"0.0.1", "0.0.2", "0.0.3", "0.0.10"}, semvers)
}

func (s *SemversTests) TestSorting_preRelease() {
	semvers := Semvers{
		Parse("1.0.0"),
		Parse("1.0.0-alpha.10"),
		Parse("1.0.0-alpha.1"),
		Parse("1.0.0-alpha"),
		Parse("1.0.0-beta.1"),
	}
	sort.Sort(semvers)
	s.assertEqual([]string{"1.0.0-alpha.1", "1.0.0-alpha.10", "1.0.0-alpha", "1.0.0-beta.1", "1.0.0"}, semvers)
}
