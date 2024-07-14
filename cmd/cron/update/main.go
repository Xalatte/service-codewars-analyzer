package main

import (
	"fmt"
	"github.com/Xalatte/service-codewars-analyzer.git/internal/usecase/html_parse"
)

func main() {
	parser := html_parse.New()
	for _, name := range parser.Names {
		fmt.Println(name, len(name))
	}
}
