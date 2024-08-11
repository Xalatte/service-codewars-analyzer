package html_parse

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNoDuplicateNames(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := os.ReadFile("test_data/codewars_leaderboard_test.html")
		_, _ = w.Write(data)
	}))
	defer svr.Close()

	ctx := context.Background()

	p := New([]string{svr.URL, svr.URL, svr.URL, svr.URL})
	names, err := p.GetUniqueLeadersNames(ctx)
	require.NoError(t, err)
	require.Equal(t, 502, len(names))
}

func TestNoDuplicateNames_ReturnError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer svr.Close()

	ctx := context.Background()

	p := New([]string{svr.URL, svr.URL, svr.URL, svr.URL})
	_, err := p.GetUniqueLeadersNames(ctx)
	require.Error(t, err)
}
