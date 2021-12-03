package main

import (
	"os"

	"github.com/nandajavarma/aoc2021/pkg/day01"
	"github.com/nandajavarma/aoc2021/pkg/day02"
	"github.com/nandajavarma/aoc2021/pkg/day03"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var inputFile string

var rootCmd = &cobra.Command{
	Use:  "aoc2021",
	Args: cobra.MinimumNArgs(1),
}

var day1 = &cobra.Command{
	Use:   "day01",
	Short: "Solution to day 01 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day01.Run(inputFile)
	},
}

var day2 = &cobra.Command{
	Use:   "day02",
	Short: "Solution to day 02 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day02.Run(inputFile)
	},
}

var day3 = &cobra.Command{
	Use:   "day03",
	Short: "Solution to day 03 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day03.Run(inputFile)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "Input file name")
	rootCmd.MarkPersistentFlagRequired("input")
}

func main() {
	rootCmd.AddCommand(day1)
	rootCmd.AddCommand(day2)
	rootCmd.AddCommand(day3)

	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("error in the cli. Exiting")
		os.Exit(1)
	}
}
