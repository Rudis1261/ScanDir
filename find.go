// Package scandir contains utilities to find a filetype in a specific location
package scandir

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var findExtension string = ""
var includeThis string = ""
var excludeThis string = ""

// Find specific file types with the filepath.Walk
func Find(params ...string) []string {

	searchDir := params[0]
	findExtension = params[1]
	found := []string{}

	// Param 3 is optional, include
	if len(params) > 2 {
		includeThis = params[2]
	}

	// Param 4 is optional, excluding
	if len(params) > 3 {
		excludeThis = params[3]
	}

	// Fill up an array with everything found
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	// Exit should we have encountered an error
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Loop through all the files we found with the walker
	// Check that it meets the criteria. Extension being the minimum
	for _, file := range fileList {
		extension := filepath.Ext(file)

		// Only allow files with an extension
		if len(extension) == 0 {
			break
		}

		// Get the file extension but trim the leading .
		extension = extension[1:len(extension)]

		// If the extensions don't match we don't need to continue testing
		if extension !== findExtension {
			continue
		}

		// Look for a need in the haystack
		if len(includeThis) > 0 && strings.Contains(file, includeThis) == false {
			continue
		}

		// Allow exclusion of files
		if len(excludeThis) > 0 && strings.Contains(file, excludeThis) {
			continue
		}

		// This must be what we are looking for, so append it to the array of strings
		found = append(found, file)
	}

	// Finally either return an empty array, or one filled with that the type of file we were looking for
	return found
}
