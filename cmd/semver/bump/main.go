package bump

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/usvc/config"
	"github.com/usvc/logger"
	"github.com/usvc/semver"
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
	hasArguments := (len(args) > 0)
	retrieveFromGit := conf.GetBool("git")
	var version *semver.Semver
	var err error

	switch true {
	case retrieveFromGit:
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to get current working directory: '%s'", err)
			os.Exit(1)
		}
		version, err = getLatestSemverFromGitRepository(currentDirectory)
		if err != nil {
			log.Errorf("failed to retrieve latest semver tag: '%s'", err)
			os.Exit(1)
		}
	case hasArguments:
		version, err = getSemverFromArguments(args)
		if err != nil {
			log.Errorf("failed to parse semver input from arguments: '%s'", err)
			os.Exit(1)
		}
	default:
		command.Help()
		return
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
		err = tagCurrentGitCommit(version.String(), currentDirectory)
		if err != nil {
			log.Errorf("failed to add tag '%s' to repository at '%s': '%s'", version.String, currentDirectory, err)
			os.Exit(1)
		}
		log.Infof("added tag '%s' to repository at '%s'", version.String, currentDirectory)
	}
}

func GetCommand() *cobra.Command {
	if cmd == nil {
		log = logger.New(logger.Options{})
		conf.LoadFromEnvironment()
		cmd = &cobra.Command{
			Use: "bump",
			Run: run,
		}
		conf.ApplyToCobra(cmd)
	}
	return cmd
}
