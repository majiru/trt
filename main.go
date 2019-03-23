package main

import (
	"bufio"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"regexp"
)

func parseFilter(name string) []string {
	var out []string
	file, err := os.Open(name)
	if err != nil {
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}

func main() {
	fp := gofeed.NewParser()
	var feed *gofeed.Feed
	var err error

	filter := parseFilter("filter.txt")
	if filter == nil {
		log.Fatal("Failed to parse filter file")
	}

	if len(os.Args) > 1 {
		feed, err = fp.ParseURL(os.Args[1])
		if err != nil {
			log.Fatal("Failed to parse URL ", os.Args[1], ":", err.Error())
		}
	} else {
		feed, err = fp.Parse(os.Stdin)
		if err != nil {
			log.Fatal("Failed to parse stdin: ", err.Error())
		}
	}

	for _, i := range feed.Items {
		for _, j := range filter {
			match, _ := regexp.MatchString(j, i.Title)
			if match == true {
				fmt.Println(i.Link)
				break
			}
		}
	}
}
