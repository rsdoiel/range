/*
 * range.go - emit a list of integers separated by spaces starting from
 * first command line parameter to last command line parameter.
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 all rights reserved.
 * Released under the Simplified BSD License
 * See: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	help      bool
	start     int
	end       int
	increment int
)

var Usage = func(exit_code int, msg string) {
	fmt.Fprintf(os.Stderr, `%s
 USAGE %s STARTING_INTEGER ENDING_INTEGER [INCREMENT_INTEGER]

 EXAMPLES
 
 Count from one through five: %s 1 5
 Count from negative two to six: %s -- -2 6
 Count even number from two to ten: %s --increment=2 2 10
 Count down from ten to one: %s 10 1

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, `

 copyright (c) 2014 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php

`)
	os.Exit(exit_code)
}

func init() {
	const (
		help_usage  = "Display this help document."
		start_usage = "The starting integer (e.g. 1)"
		end_usage   = "The ending integer (e.g. 10)"
		inc_usage   = "The non-zero integer value to increment by (e.g. 1 or -1)"
	)

	flag.IntVar(&start, "start", 0, start_usage)
	flag.IntVar(&start, "s", 0, start_usage)
	flag.IntVar(&end, "end", 9, end_usage)
	flag.IntVar(&end, "e", 9, end_usage)
	flag.IntVar(&increment, "increment", 1, inc_usage)
	flag.IntVar(&increment, "i", 1, inc_usage)

	flag.BoolVar(&help, "help", help, help_usage)
	flag.BoolVar(&help, "h", help, help_usage)
}

func assertOk(e error, fail_msg string) {
	if e != nil {
		Usage(1, fmt.Sprintf(" %s\n %s\n", fail_msg, e))
	}
}

func main() {
	flag.Parse()
	if help == true {
		Usage(0, "")
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 2 {
		Usage(1, "Must include start and end integers.")
	} else if argc > 3 {
		Usage(1, "Too many command line arguments.")
	}

	start, err := strconv.Atoi(argv[0])
	assertOk(err, "Start value must be an integer.")
	end, err := strconv.Atoi(argv[1])
	assertOk(err, "End value must be an integer.")
	if argc == 3 {
		increment, err = strconv.Atoi(argv[2])
	} else if increment == 0 {
		err = errors.New("increment was zero.")
	}
	assertOk(err, "Increment must be a non-zero integer.")

	if (end < start) && (increment > 0) {
		increment = increment * -1
	} else if (start < end) && (increment < 0) {
		increment = increment * -1
	}
	for i := start; i <= end; i = i + increment {
		if i == start {
			fmt.Printf("%d", i)
		} else {
			fmt.Printf(" %d", i)
		}
	}
}
