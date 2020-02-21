package semver

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BuildMetadataTests struct {
	suite.Suite
}

func TestBuildMetadata(t *testing.T) {
	suite.Run(t, &BuildMetadataTests{})
}

func (s *BuildMetadataTests) TestString() {
	buildMetadata := BuildMetadata{"a", "b", "c"}
	s.Equal("a.b.c", buildMetadata.String())
}
