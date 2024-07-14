package html_parse

import (
	"testing"
)

func TestCodeWarsLeaderBoard(t *testing.T) {

}
func TestNoDuplicateNames(t *testing.T) {
	p := New()
	namesMap := make(map[string]struct{})
	for _, name := range p.Names {
		if _, exists := namesMap[name]; exists {
			t.Errorf("найден повторяющийся ник: %s", name)
		}
		namesMap[name] = struct{}{}
	}
}
