package bump

import (
	"github.com/spf13/cobra"
	"github.com/usvc/logger"
)

var (
	cmd *cobra.Command
	log logger.Logger
)

func main(command *cobra.Command, args []string) {
	log.Info("hello")
}

func GetCommand() *cobra.Command {
	if cmd == nil {
		cmd = &cobra.Command{
			Use: "bump",
			Run: main,
		}
	}
	return cmd
}
