package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "A collection of advent of code challanges",
	Long: `Advent of code 2024 made as a CLI tool
	to make it easier to manage and run each of the challanges for the day`}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
