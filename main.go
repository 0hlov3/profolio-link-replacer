package main

import (
	"fmt"
	"os"
	"regexp"
)

// Fallback/default URLs
var defaultLinks = map[string]string{
	"NO_LINK": "#",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path-to-markdown-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Println("File path:", filePath)

	// Read file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)

	// Match both [label](<url>) and [label](<url> "title")
	re := regexp.MustCompile(`\[(.*?)\]\(<(.*?)>(?:\s+"(.*?)")?\)`)

	// Replace function
	newContent := re.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		label := submatches[1]
		url := submatches[2]
		title := ""
		if len(submatches) >= 4 {
			title = submatches[3]
		}

		if url == "" {
			if fallbackURL, ok := defaultLinks[title]; ok {
				url = fallbackURL
			}
		}

		isBold := false
		if matched, _ := regexp.MatchString(`^\*\*(.*)\*\*$`, label); matched {
			isBold = true
			label = label[2 : len(label)-2] // strip ** for inner content
		}

		replacement := fmt.Sprintf(`{{< newtablink "%s" >}}%s{{< /newtablink >}}`, url, label)
		if isBold {
			replacement = fmt.Sprintf("**%s**", replacement)
		}

		return replacement
	})

	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Links updated in:", filePath)
}
