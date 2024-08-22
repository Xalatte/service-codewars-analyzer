package kata_ids

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Xalatte/service-codewars-analyzer.git/internal/usecase/html_parse"
	"io"
	"net/http"
)

type Usecase struct {
	TotalPages int    `json:"totalPages"`
	KataIds    string `json:"id"`
	name       string
}

func New() *Usecase {
	return &Usecase{}
}
func (u *Usecase) GetKataID(id int) ([]string, error) {
	return nil, nil
}

func (u *Usecase) GetHttpsLen(context.Context, *http.Client) (int, string, error) {
	ctx := context.Background()
	parser := html_parse.New([]string{
		html_parse.KataURL,
		html_parse.LeadersURL,
		html_parse.AuthoredURL,
		html_parse.RanksURL,
	})

	names, err := parser.GetUniqueLeadersNames(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, name := range names {
		resp, err := http.Get("http://www.codewars.com/api/v1/users/" + name + "/code-challenges/completed?page=0")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var usecase Usecase
		u.name = name
		if err := json.Unmarshal(body, &usecase); err != nil {
			fmt.Println(err)
		}
	}
	return u.TotalPages, u.name, nil
}
