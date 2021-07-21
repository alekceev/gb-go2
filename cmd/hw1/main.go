package main

import (
	// "gb-go2/pkg/hw1"
	"log"

	search "github.com/alekceev/gosearcher/v2"
)

func main() {
	urls := []string{
		"https://restcountries.eu/rest/v2/alpha/ru",
		"https://restcountries.eu/rest/v2/alpha/us",
		"https://restcountries.eu/rest/v2/alpha/es",
		"https://restcountries.eu/rest/v2/alpha/au",
	}

	find := "English"

	res, err := search.SearchInUrls(find, urls)

	if err != nil {
		log.Fatal(err)
	}

	for _, u := range res {
		log.Printf("found %s in %s\n", find, u)
	}
}
