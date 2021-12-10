package main

import (
	"os"

	"github.com/nandajavarma/aoc2021/pkg/day01"
	"github.com/nandajavarma/aoc2021/pkg/day02"
	"github.com/nandajavarma/aoc2021/pkg/day03"
	"github.com/nandajavarma/aoc2021/pkg/day04"
	"github.com/nandajavarma/aoc2021/pkg/day05"
	"github.com/nandajavarma/aoc2021/pkg/day06"
	"github.com/nandajavarma/aoc2021/pkg/day07"
	"github.com/nandajavarma/aoc2021/pkg/day08"
	"github.com/nandajavarma/aoc2021/pkg/day09"
	"github.com/nandajavarma/aoc2021/pkg/day10"
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

var day4 = &cobra.Command{
	Use:   "day04",
	Short: "Solution to day 04 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day04.Run(inputFile)
	},
}

var day5 = &cobra.Command{
	Use:   "day05",
	Short: "Solution to day 05 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day05.Run(inputFile)
	},
}

var day6 = &cobra.Command{
	Use:   "day06",
	Short: "Solution to day 06 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day06.Run(inputFile)
	},
}

var day7 = &cobra.Command{
	Use:   "day07",
	Short: "Solution to day 07 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day07.Run(inputFile)
	},
}

var day8 = &cobra.Command{
	Use:   "day08",
	Short: "Solution to day 08 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day08.Run(inputFile)
	},
}

var day9 = &cobra.Command{
	Use:   "day09",
	Short: "Solution to day 09 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day09.Run(inputFile)
	},
}

var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "Solution to day 10 of AoC 2021",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return day10.Run(inputFile)
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
	rootCmd.AddCommand(day4)
	rootCmd.AddCommand(day5)
	rootCmd.AddCommand(day6)
	rootCmd.AddCommand(day7)
	rootCmd.AddCommand(day8)
	rootCmd.AddCommand(day9)
	rootCmd.AddCommand(day10Cmd)

	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("error in the cli. Exiting")
		os.Exit(1)
	}
}
