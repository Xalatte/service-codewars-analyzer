package html_parse

import (
	"context"
	"testing"
)

func TestGetNamesLeaders(t *testing.T) {
	tCase := []struct {
		name        string
		expectedLen int
	}{
		{
			name:        "g964",
			expectedLen: 511,
		},
	}

	for _, tc := range tCase {
		t.Run(tc.name, func(t *testing.T) {
			parser := New()
			names, err := parser.GetNamesLeaders(context.Background())
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(names) != tc.expectedLen {
				t.Errorf("Expected %d names, but got %d", tc.expectedLen, len(names))
			}
		})
	}
}
