package main

import (
	"context"
	"fmt"
	"github.com/Xalatte/service-codewars-analyzer.git/internal/usecase/leaderboards_names"
	"log"
)

func main() {
	ctx := context.Background()
	parser := leaderboards_names.New([]string{
		leaderboards_names.KataURL,
		leaderboards_names.LeadersURL,
		leaderboards_names.AuthoredURL,
		leaderboards_names.RanksURL,
	})

	names, err := parser.GetUniqueLeadersNames(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
