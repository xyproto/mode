package mode

import (
	"fmt"
	"testing"
)

// describe returns a readable description of ts, since its String method
// returns the indentation itself
func describe(ts TabsSpaces) string {
	return fmt.Sprintf("{PerTab: %d, Spaces: %v}", ts.PerTab, ts.Spaces)
}

func TestTabsSpaces(t *testing.T) {
	for m, want := range map[Mode]TabsSpaces{
		YAML:      {2, true},
		JSON:      {2, true},
		Go:        {4, false},
		Python:    {4, true},
		Fish:      {4, true},
		Oil:       {2, true},
		Fortran77: {7, true},
		Blank:     DefaultTabsSpaces,
	} {
		if got := m.TabsSpaces(); got != want {
			t.Errorf("%s: expected %s, got %s", m, describe(want), describe(got))
		}
		if got := m.Spaces(); got != want.Spaces {
			t.Errorf("%s: expected Spaces() to be %v, got %v", m, want.Spaces, got)
		}
	}
}

// A mode that is listed for more than one indentation setting gets a random
// one of them, since map iteration order is random
func TestEachModeHasOneIndentation(t *testing.T) {
	seen := make(map[Mode]TabsSpaces)
	for ts, modes := range languageIndentation {
		for _, m := range modes {
			if prev, ok := seen[m]; ok && prev != ts {
				t.Errorf("%s is listed for both %s and %s", m, describe(prev), describe(ts))
			}
			seen[m] = ts
		}
	}
}
