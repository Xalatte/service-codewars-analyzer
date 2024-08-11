package main

import (
	"context"
	"fmt"
	"github.com/Xalatte/service-codewars-analyzer.git/internal/usecase/html_parse"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	parser := html_parse.New([]string{
		html_parse.KataURL,
		html_parse.LeadersURL,
		html_parse.AuthoredURL,
		html_parse.RanksURL,
	})

	startTime := time.Now()
	_, err := parser.GetUniqueLeadersNames(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(startTime).String())
}
