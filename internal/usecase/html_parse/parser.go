package html_parse

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

// GetNamesLeaders возращает список уникальных ников из таблиц лидеров
func (p Parser) GetNamesLeaders(ctx context.Context) ([]string, error) {
	c := colly.NewCollector()

	names := make(map[string]bool)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		parts := strings.Split(href, "/")
		if len(parts) > 2 {
			names[parts[2]] = true
		}
	})

	c.Visit("https://www.codewars.com/users/leaderboard")

	c.Wait()

	result := make([]string, 0, len(names))
	for name := range names {
		result = append(result, name)
	}

	return result, nil
}

func main() {
	parser := New()
	names, err := parser.GetNamesLeaders(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
