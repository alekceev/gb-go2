package main

import (
	"gb-go2/pkg/hw1"
	"log"
)

func main() {
	urls := []string{
		"https://restcountries.eu/rest/v2/alpha/ru",
		"https://restcountries.eu/rest/v2/alpha/us",
		"https://restcountries.eu/rest/v2/alpha/es",
		"https://restcountries.eu/rest/v2/alpha/au",
	}

	search := "English"

	res, err := hw1.SearchInUrls(search, urls)

	if err != nil {
		log.Fatal(err)
	}

	for _, u := range res {
		log.Printf("found %s in %s\n", search, u)
	}
}
