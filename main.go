package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// Fallback/default URLs
var defaultLinks = map[string]string{
	"": "",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path-to-markdown-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Println("File path:", filePath)

	// Read file
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)

	// Regex: [Label](<URL> "Title")
	re := regexp.MustCompile(`\[(.*?)\]\(<(.*?)> "(.*?)"\)`)

	// Replace function
	newContent := re.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) != 4 {
			return match // fallback if regex breaks
		}

		label := submatches[1]
		url := submatches[2]
		title := submatches[3]

		if url == "" {
			if fallbackURL, ok := defaultLinks[title]; ok {
				url = fallbackURL
			}
		}

		return fmt.Sprintf(`{{< newtablink "%s" >}}%s{{< /newtablink >}}`, url, label)
	})

	// Write back in place
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Links updated in:", filePath)
}
