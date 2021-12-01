# Advent of Code 2021

To use this module, clone this repo and build the module using:

``` sh
$ go build -o aoc2021 ./cmd/aosc2021/main.go
```

You will find a binary in the current working directory named `aoc2021`.

Try running it by doing:

``` sh
$ ./aoc2021 -h
Usage:
  aoc2021 [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  day01       Solution to day 01 of AoC 2021
  help        Help about any command

Flags:
  -h, --help           help for aoc2021
      --input string   Input file name
```

Like the usage suggests you can choose an available command with the mandatory
input file flag. For example, to see the solution of day 1, run the command:

``` sh
$ ./aoc2021 day01 --input ./assets/day01/input
```
