package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	// Define a flag to accept URLs as arguments
	flag.String("urls", "", "URLs to add")

	// Parse the flags
	flag.Parse()

	// Get the URLs from the flag
	urls := flag.Args()

	// Compile regular expressions to match Instagram and Twitter URLs
	instagramRegex, err := regexp.Compile("^(https?)://(www.)?instagram.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	twitterRegex, err := regexp.Compile("^(https?)://(www.)?twitter.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Filter the URLs to only allow Instagram and Twitter URLs
	var filteredURLs []string
	for _, url := range urls {
		if instagramRegex.MatchString(url) || twitterRegex.MatchString(url) {
			filteredURLs = append(filteredURLs, url)
		}
	}

	// Read the contents of the urls.json file
	data, err := ioutil.ReadFile("urls.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmarshal the JSON data
	var existingURLs []string
	if err := json.Unmarshal(data, &existingURLs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Append the new URLs to the existing URLs
	existingURLs = append(existingURLs, filteredURLs...)

	// Marshal the updated URLs
	updatedURLs, err := json.Marshal(existingURLs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write the updated URLs to the urls.json file
	if err := ioutil.WriteFile("urls.json", updatedURLs, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
