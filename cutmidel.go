package main

import (
	"fmt"
	"os"
	"strconv"
)

const version = "1.0.0"

const hint = "use -h to get help"

func usage() {
	fmt.Printf(`cutmidel %s
  Shortens a text by trimming it by an ellipsis in the middle.

Usage: cutmidel [options] <text> <leading> <trailing> [ellipsis]

  Specify the text and the maximum count of leading and trailing
  characters. The overall maximum length will be their sum plus
  the length of an ellipsis (3 dots by default). Zero for either
  leading or trailing count means no leading or trailing parts.

Options:
  -V|--version  prints the version of the executable and exits
  -h|--help     prints th usage information and exits

Examples:
  $ cutmidel "~/Sources/private/cutmidel" 5 10
  ~/Sou...e/cutmidel
  $ cutmidel ~/Sources/private/cutmidel 0 12 ..
  ..ate/cutmidel
`, version)
}

func main() {
	argc := len(os.Args)
	// print usage information if no arguments were provided
	if argc == 1 {
		usage()
		os.Exit(1)
	}
	// check if printing the version number or usage information was requested
	if argc == 2 {
		arg := os.Args[1]
		if arg == "-V" || arg == "--version" {
			println(version)
			os.Exit(0)
		}
		if arg == "-h" || arg == "--help" {
			usage()
			os.Exit(0)
		}
		// unexpected one argument or fail if unexpected arguments were not provided
		fmt.Printf("invalid argument: %s (%s)\n", arg, hint)
		os.Exit(1)
	}
	// check if exactly three arguments are available
	if argc == 3 {
		fmt.Printf("too few arguments (%s)\n", hint)
		os.Exit(1)
	}
	if argc > 5 {
		fmt.Printf("too many arguments (%s)\n", hint)
		os.Exit(1)
	}

	// make sure that leading and trailing character count are numeric
	var err error
	var lead int
	if lead, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Printf("invalid leading character count: \"%s\" (%s)\n", os.Args[2], hint)
		os.Exit(1)
	}
	var trail int
	if trail, err = strconv.Atoi(os.Args[3]); err != nil {
		fmt.Printf("invalid trailing character count: \"%s\" (%s)\n", os.Args[3], hint)
		os.Exit(1)
	}
	// ellipsis cannot be put to the middle unless the middle is specified
	if lead == 0 && trail == 0 {
		fmt.Printf("both leading and trailing counts cannot be zero (%s)\n", hint)
		os.Exit(1)
	}
	// check if a custom ellipsis was specified
	var ellipsis string
	if argc == 5 {
		ellipsis = os.Args[4]
	} else {
		ellipsis = "..."
	}
	elliplen := len(ellipsis)

	// get the text to trim with its length; it will be trimmed in-place
	txt := os.Args[1]
	txtlen := len(txt)

	// if the input text is shorter than the ledting and trailing character
	// count plus the ellipsis length, leave it intact
	if txtlen > lead+trail+elliplen {
		if lead == 0 {
			// if only the trailing part should be displayed, put the ellipsis
			// to the beginning and move the trailing part behind it
			txt = ellipsis + txt[txtlen-trail:]
		} else if trail == 0 {
			// if only the leading part should be displayed, put the ellipsis
			// right behind it
			txt = txt[:lead] + ellipsis
		} else {
			// if both leading and trailing parts should be displayed, put the
			// ellipsis behind the leading part and move the trailing part behind it
			txt = txt[:lead] + ellipsis + txt[txtlen-trail:]
		}
	}

	print(txt)
}
