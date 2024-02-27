// todotxt prints out tasks by project from todo.txt formatted file
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

func main() {

	todoFile := flag.String("file", "", "Only search this file for todo lines.")
	flag.Parse()

	var rootPath string = os.Getenv(`todotxt`)

	projectRegex, _ := regexp.Compile(`(\+\W+$)`) // How we find projects
	projects := make(map[string][]string)

	if *todoFile != "" {
		rootPath = *todoFile
	}
	walkErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Unable to read file: %s - %s", path, err)
		}

		scanner := bufio.NewScanner(file)
		var keys []string;
		for scanner.Scan() {
			var todoLine string = scanner.Text()

			// warning: only supporting one project per line. 
			// Pull requests welcome.
			keys = projectRegex.FindAllString(todoLine, -1)
			for _, key := range keys{
				if _, ok := projects[key]; ok {
					projects[key] = append(projects[key], todoLine)
				} else {
					projects[key] = []string{todoLine}
				}
			}

		}

		for key, strings := range projects{
			fmt.Printf("##: %s\n", key)
			for _, str := range strings {
				fmt.Printf("  - %s\n", str)
			}
		}

		return nil
	})

	if walkErr != nil {
		log.Fatal(walkErr)
	}

}
