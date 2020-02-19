package main

import (
	"bufio"
	"fmt"
	"hw3_bench/json_struct"
	"io"
	"os"
	"strings"
)


func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	seenBrowsers := make([]string, 0, 100)
	var isAndroid, isMSIE, notSeenBefore bool
	fmt.Fprintln(out, "found users:")
	for i := 0; scanner.Scan(); i++ {
		user := json_struct.User{}
		_ = user.UnmarshalJSON([]byte(scanner.Text()))
		isAndroid = false
		isMSIE = false
		for _, browser := range user.Browsers {
			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				notSeenBefore = true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
						break
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
				}
			} else if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				notSeenBefore = true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
						break
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
				}
			}
		}

		if !isAndroid || !isMSIE {
			continue
		}
		fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", i, user.Name, strings.ReplaceAll(user.Email, "@", " [at] ")))
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
