package scrapper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//TODO: add function to get id from title

func ScrapeWebsite(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("failed to fetch url: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	//TODO: add logic to get data of anime
	title := doc.Find("title").Text()
	return title, nil
}
