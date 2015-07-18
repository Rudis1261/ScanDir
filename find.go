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

// Using the directory walker. We will scan the directory for a specific file extension
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

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, file := range fileList {
		extension := filepath.Ext(file)

		// Only allow files with an extension
		if len(extension) > 0 {

			// Trim the leading .
			extension = extension[1:len(extension)]

			if extension == findExtension {

				// Look for a need in the haystack
				if len(includeThis) > 0 && strings.Contains(file, includeThis) == false {
					continue
				}

				// Allow exclusion of files
				if len(excludeThis) > 0 && strings.Contains(file, excludeThis) {
					continue
				}

				found = append(found, file)
			}
		}
	}
	return found
}
