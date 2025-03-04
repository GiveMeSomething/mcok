package mcok

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var outputPath string = "result.txt"

var MockCommand = &cobra.Command{
	Use:   "mock",
	Short: "Start data generation",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		err := writeOutput(outputPath, 1_000)
		if err != nil {
			fmt.Printf("Failed: %s\n", err)
			return
		}

		fmt.Printf("Generation finished: %s\n", time.Since(start))
	},
}
