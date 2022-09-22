package app

import (
	"fmt"
	"runtime"

	"simpleagent/pkg/version"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	AgentCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if flagNoColor {
			color.NoColor = true
		}
		av, _ := version.Agent()
		meta := ""
		if av.Meta != "" {
			meta = fmt.Sprintf("- Meta: %s ", color.YellowString(av.Meta))
		}
		fmt.Fprintln(
			color.Output,
			fmt.Sprintf("Agent %s %s- Commit: %s - Go version: %s",
				color.CyanString(av.GetNumberAndPre()),
				meta,
				color.GreenString(av.Commit),
				color.RedString(runtime.Version()),
			),
		)
	},
}
