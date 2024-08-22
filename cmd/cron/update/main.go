package main

import (
	"context"
	"fmt"
	"github.com/Xalatte/service-codewars-analyzer.git/internal/usecase/html_parse"
	"log"
)

func main() {
	ctx := context.Background()
	parser := html_parse.New([]string{
		html_parse.KataURL,
		html_parse.LeadersURL,
		html_parse.AuthoredURL,
		html_parse.RanksURL,
	})

	names, err := parser.GetUniqueLeadersNames(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
