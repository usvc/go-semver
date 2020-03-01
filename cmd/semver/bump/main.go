package bump

import (
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/config"
	"github.com/usvc/logger"
	"github.com/usvc/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	cmd  *cobra.Command
	log  logger.Logger
	conf = config.Map{
		"major":       &config.Bool{Shorthand: "M"},
		"minor":       &config.Bool{Shorthand: "m"},
		"patch":       &config.Bool{Shorthand: "p"},
		"pre-release": &config.Bool{Shorthand: "P"},
		"git":         &config.Bool{Shorthand: "g"},
		"apply":       &config.Bool{Shorthand: "A"},
	}
)

func run(command *cobra.Command, args []string) {
	var version *semver.Semver
	if conf.GetBool("git") {
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to getting current working directory: '%s'", err)
			os.Exit(1)
		}
		repository, err := git.PlainOpen(currentDirectory)
		if err != nil {
			log.Errorf("failed to open the directory '%s' as a git repository: '%s'", currentDirectory, err)
			os.Exit(1)
		}
		tags, err := repository.Tags()
		if err != nil {
			log.Errorf("failed to get tags from git repository at '%s': '%s'", currentDirectory, err)
			os.Exit(1)
		}
		var versions semver.Semvers
		tags.ForEach(func(tag *plumbing.Reference) error {
			tagName := tag.Name().Short()
			if semver.IsValid(tagName) {
				versions = append(versions, semver.Parse(tagName))
			}
			return nil
		})
		sort.Sort(versions)
		log.Tracef("retrieved tags: %v", versions)
		version = versions[versions.Len()-1]
		log.Debugf("using version from git tags: '%s'", version)
	} else if len(args) == 0 {
		command.Help()
		return
	} else {
		versionFromArguments := strings.Join(args, ".")
		if !semver.IsValid(versionFromArguments) {
			log.Errorf("parsed input '%s' is not a valid semver string", versionFromArguments)
			os.Exit(1)
		}
		version = semver.Parse(versionFromArguments)
		log.Debugf("using version from provided input: '%s'", version)
	}
	log.Debugf("current version: '%s'", version)
	switch true {
	case conf.GetBool("major"):
		version.BumpMajor()
	case conf.GetBool("minor"):
		version.BumpMinor()
	case conf.GetBool("pre-release"):
		version.BumpPatch()
	case conf.GetBool("patch"):
		fallthrough
	default:
		version.BumpPatch()
	}
	log.Debugf("next version: '%s'", version)

	if conf.GetBool("git") && conf.GetBool("apply") {
		log.Debugf("adding git tag '%s' to repository...", version.String())

		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to getting current working directory: '%s'", err)
			os.Exit(1)
		}
		repository, err := git.PlainOpen(currentDirectory)
		if err != nil {
			log.Errorf("failed to open the directory '%s' as a git repository: '%s'", currentDirectory, err)
			os.Exit(1)
		}

		head, err := repository.Head()
		if err != nil {
			log.Errorf("failed to get HEAD of git repository: '%s'", err)
			os.Exit(1)
		}
		headCommitHash := head.Hash()
		repository.CreateTag(version.String(), headCommitHash, nil)

		log.Infof("added git tag '%s' to repository at HEAD '%s'", version.String(), headCommitHash.String())
	}
}

func GetCommand() *cobra.Command {
	if cmd == nil {
		initialize()
	}
	return cmd
}

func initialize() {
	log = logger.New(logger.Options{})
	conf.LoadFromEnvironment()
	cmd = &cobra.Command{
		Use: "bump",
		Run: run,
	}
	conf.ApplyToCobra(cmd)
}
