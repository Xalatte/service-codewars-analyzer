package html_parse

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

const (
	kataURL     = "https://www.codewars.com/users/leaderboard/kata"
	authoredURL = "https://www.codewars.com/users/leaderboard/authored"
	ranksURL    = "https://www.codewars.com/users/leaderboard/ranks"
	leadersURL  = "https://www.codewars.com/users/leaderboard"
)

type Parser struct {
	urls  []string
	Names []string
}

func New() *Parser {
	constants := []string{kataURL, authoredURL, ranksURL, leadersURL}
	p := &Parser{urls: constants}
	ch := make(chan []string)

	// Используем множество для хранения уникальных ников
	namesSet := make(map[string]struct{})

	for _, url := range constants {
		go func(url string) {
			names, err := p.getNamesLeaders(context.Background(), url)
			if err != nil {
				fmt.Printf("Error fetching names from %s: %v\n", url, err)
				ch <- nil
				return
			}
			ch <- names
		}(url)
	}

	for range constants {
		names := <-ch
		if names != nil {
			for _, name := range names {
				if _, exists := namesSet[name]; !exists {
					namesSet[name] = struct{}{}
					p.Names = append(p.Names, name)
				}
			}
		}
	}

	fmt.Println("All URLs processed")
	return p
}

func (p Parser) getNamesLeaders(_ context.Context, url string) ([]string, error) {
	c := colly.NewCollector()

	namesSet := make(map[string]struct{})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		parts := strings.Split(href, "/")
		if len(parts) > 2 && parts[1] == "users" {
			namesSet[parts[2]] = struct{}{}
		}
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("c.Visit: %w", err)
	}

	names := make([]string, 0, len(namesSet))
	for name := range namesSet {
		names = append(names, name)
	}

	return names, nil
}
