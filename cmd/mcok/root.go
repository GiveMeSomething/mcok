package mcok

import (
	"fmt"
	"givemesomething/mcok/internal/mcok/controller"
	"givemesomething/mcok/internal/mcok/generator"
	"time"

	"github.com/spf13/cobra"
)

var outputPath string
var count int64

var MockCommand = &cobra.Command{
	Use:   "mock",
	Short: "Start data generation",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		err := controller.CreateOutputDirIfNotExist(outputPath)
		if err != nil {
			fmt.Printf("Failed: %s\n", err)
			return
		}

		err = controller.WriteOutput(
			outputPath,
			count,
			func() string {
				return generator.RandomPassenger().ToString()
			},
		)
		if err != nil {
			fmt.Printf("Failed: %s\n", err)
			return
		}

		fmt.Printf("Generation finished: %s\n", time.Since(start))
	},
}

func init() {
	MockCommand.Flags().StringVarP(&outputPath, "output", "o", "output/result.txt", "Define output path. Default to output/result.txt")
	MockCommand.Flags().Int64VarP(&count, "count", "c", 100, "Number of data rows needed. Default to 100")
}
