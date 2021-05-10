package df2discord

import (
	"regexp"
)

var colorFilter = regexp.MustCompile("(§.)")

// filterColor removes all minecraft color codes.
// This includes invalid color codes, as those do not show up in-game.
func filterColor(s string) string {
	return colorFilter.ReplaceAllString(s, "")
}
