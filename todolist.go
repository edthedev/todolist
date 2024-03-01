// todolist finds lines starting with `- [ ]` in `.md` files.
// It expects a `todolist` environment variable to tell it where to look for
// files.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func searchForMatchesByLine(tagRegex regexp.Regexp, fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Unable to read file: %s - %s", fileName, err)
	}
	scanner := bufio.NewScanner(file)
	var results = []string{}
	var result = []string{}
	for scanner.Scan() {
		result = tagRegex.FindAllString(scanner.Text(), -1)
		results = append(results, result...)
	}

	return results
}

func main() {

	todoFile := flag.String("file", "", "Only search this file for todo lines.")
	// doneFlag := flag.Bool("done", false, "List any done lines.")
	// maxFlag := flag.Int("max", 5, "Maximum number of files or tags to list.")
	flag.Parse()

	todoRegex, _ := regexp.Compile(`(^[\-\+]\W\[\W+\].+$)`) // How we find todos
	// doneRegex, _ := regexp.Compile(`\[x\]|\[\/\]`)      // How we find done items.

	var rootPath string = os.Getenv(`todolist`)
	var notMarkdown int = 0
	// var matchCount int = 0

	// fmt.Printf("List: %t, Find: %s.", *listFlag, *searchFlag)
	// fmt.Printf("List: %t, Find: %s.", "--default to list--", *searchFlag)
	// fmt.Println("")

	if *todoFile != "" {
		rootPath = *todoFile
	}
	walkErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) != ".md" {
			notMarkdown += 1
			return nil
		}

		var found = []string{}
		found = searchForMatchesByLine(*todoRegex, path)
		// fmt.Println(found)
		var todo string

		if len(found) > 0 {
			fmt.Println("")
			fmt.Println("## ", path)
		}
		for _, todo = range found {
			fmt.Println(todo)
			// fmt.Println("")
		}

		/*
			if found {
				// fmt.Printf("Found a match in: %s", path)
				matchCount += 1
				results = append(results, path)
			} else {
				// fmt.Printf("Found no match in: %s", path)
			}
		*/

		return nil
	})

	if walkErr != nil {
		log.Fatal(walkErr)
	}

}
