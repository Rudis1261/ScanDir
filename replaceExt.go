// Package scandir contains utilities to find a filetype in a specific location
package scandir

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var input string = ""
var fromExt string = ""
var toExt string = ""

// Replace the file extension in a string with another
func ReplaceExt(params ...string) string {

	// Break when too few params are provided
	if len(params) <= 2 {
		log.Println("Missing required parameters")
		log.Println("scandir.ReplaceExt(\"<input string>\", \"<from extension>\", \"<to extension>\")")
		log.Println("scandir.ReplaceExt(\"/home/user/directory/file.css\", \".css\", \".min.css\")")
		os.Exit(1)
	}

	// Param 1, The string
	if len(params) > 0 {
		input = params[0]
	}

	// Param 2, From file extension
	if len(params) > 1 {
		fromExt = params[1]
	}

	// Param 3, To extension
	if len(params) > 2 {
		toExt = params[2]
	}

	// If the input string did have a file extension. We can get it
	// And replace it with the extension we want to get out
	currentExt := filepath.Ext(input)
	if len(currentExt) > 0 {
		input = strings.TrimSuffix(input, currentExt) + toExt
	}

	return input
}
