package bump

import (
	"fmt"
	"sort"
	"strings"

	"github.com/usvc/go-semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func tagCurrentGitCommit(withTag string, pathToRepository string) error {
	repository, err := git.PlainOpen(pathToRepository)
	if err != nil {
		return fmt.Errorf("failed to open the directory '%s' as a git repository: '%s'", pathToRepository, err)
	}
	head, err := repository.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD of git repository: '%s'", err)
	}
	headCommitHash := head.Hash()
	_, err = repository.CreateTag(withTag, headCommitHash, nil)
	return err
}

func getLatestSemverFromGitRepository(pathToRepository string) (*semver.Semver, error) {
	repository, err := git.PlainOpen(pathToRepository)
	if err != nil {
		return nil, fmt.Errorf("failed to open the directory '%s' as a git repository: '%s'", pathToRepository, err)
	}
	err = repository.Fetch(&git.FetchOptions{})
	if err != nil {
		if err != git.NoErrAlreadyUpToDate {
			return nil, fmt.Errorf("failed to fetch information from the remote: '%s'", err)
		}
	}
	tags, err := repository.Tags()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags from git repository at '%s': '%s'", pathToRepository, err)
	}
	var versions semver.Semvers
	tags.ForEach(func(tag *plumbing.Reference) error {
		tagName := tag.Name().Short()
		if semver.IsValid(tagName) {
			versions = append(versions, semver.Parse(tagName))
		}
		return nil
	})
	if versions.Len() == 0 {
		return nil, fmt.Errorf("failed to retrieve any semver tags from git repository at '%s'", pathToRepository)
	}
	sort.Sort(versions)
	return versions[versions.Len()-1], nil
}

func getSemverFromArguments(args []string) (*semver.Semver, error) {
	versionFromArguments := strings.Join(args, ".")
	if !semver.IsValid(versionFromArguments) {
		return nil, fmt.Errorf("parsed input '%s' is not a valid semver string", versionFromArguments)
	}
	return semver.Parse(versionFromArguments), nil
}
