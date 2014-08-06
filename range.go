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
	var fh = os.Stderr

	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s STARTING_INTEGER ENDING_INTEGER [INCREMENT_INTEGER]

 EXAMPLES
 
 Count from one through five: %s 1 5
 Count from negative two to six: %s -- -2 6
 Count even number from two to ten: %s --increment=2 2 10
 Count down from ten to one: %s 10 1

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
	})

	fmt.Fprintf(fh, `

 copyright (c) 2014 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php

`)
	os.Exit(exit_code)
}

func init() {
	const (
		help_usage  = "Display this help document."
		start_usage = "The starting integer."
		end_usage   = "The ending integer."
		inc_usage   = "The non-zero integer increment value."
	)

	flag.IntVar(&start, "start", 0, start_usage)
	flag.IntVar(&start, "s", 0, start_usage)
	flag.IntVar(&end, "end", 0, end_usage)
	flag.IntVar(&end, "e", 0, end_usage)
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

func inRange(i, start, end int) bool {
    if start <= end && i <= end {
        return true
    }
    if start >= end && i >= end {
        return true
    }
    return false
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

    if start == end {
      fmt.Printf("%d", start)
      os.Exit(0)
    }

	// Normalize to a positive value.
	if start <= end && increment < 0 {
		increment = increment * -1
	}
	if start > end && increment > 0 {
		increment = increment * -1
	}

	// Now count up or down as appropriate.
	for i := start; inRange(i, start, end) == true; i = i + increment {
		if i == start {
			fmt.Printf("%d", i)
		} else {
			fmt.Printf(" %d", i)
		}
	}
}
