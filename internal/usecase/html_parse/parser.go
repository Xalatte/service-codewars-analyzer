package html_parse

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
)

const (
	KataURL     = "https://www.codewars.com/users/leaderboard/kata"
	AuthoredURL = "https://www.codewars.com/users/leaderboard/authored"
	RanksURL    = "https://www.codewars.com/users/leaderboard/ranks"
	LeadersURL  = "https://www.codewars.com/users/leaderboard"
)

type Parser struct {
	urls []string
}

func New(urls []string) *Parser {
	return &Parser{
		urls: urls,
	}
}

func (p *Parser) GetUniqueLeadersNames(ctx context.Context) ([]string, error) {
	errGroup, grCtx := errgroup.WithContext(ctx)
	var mut sync.Mutex

	namesMap := map[string]struct{}{}

	for _, url := range p.urls {
		errGroup.Go(func() error {
			names, err := p.getNamesLeaders(grCtx, url)
			if err != nil {
				return fmt.Errorf("p.getNamesLeaders: %w", err)
			}
			mut.Lock()
			for _, name := range names {
				namesMap[name] = struct{}{}
			}
			mut.Unlock()
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return nil, fmt.Errorf("failed get names: %w", err)
	}

	result := make([]string, 0, len(namesMap))
	for name, _ := range namesMap {
		result = append(result, name)
	}
	return result, nil
}

func (p *Parser) getNamesLeaders(_ context.Context, url string) ([]string, error) {
	c := colly.NewCollector()
	names := make([]string, 0)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		parts := strings.Split(href, "/")
		if len(parts) > 2 && parts[1] == "users" {
			names = append(names, parts[2])
		}
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("c.Visit: %w", err)
	}
	return names, nil
}
