package grep

import (
	"fmt"
	"regexp"
	"strings"
)

func (app *App) filter(r *regexp.Regexp) string {
	matches := app.findMatches(r)

	if app.count {
		if len(matches) == 0 {
			return "No matches found"
		} else {
			return fmt.Sprintf("Matches found: %v", len(matches))
		}
	}

	res := strings.Builder{}

	for _, v := range matches {
		for j := 1; j < app.before+1; j++ {
			if v-j >= 0 {
				if app.printLineNum {
					res.WriteString(fmt.Sprintf("%v-%v\n", v-j+1, app.input[v-j]))
				} else {
					res.WriteString(fmt.Sprintf("%v\n", app.input[v-j]))
				}
			}
		}

		if app.printLineNum {
			res.WriteString(fmt.Sprintf("%v:%v\n", v+1, app.input[v]))
		} else {
			res.WriteString(fmt.Sprintf("%v\n", app.input[v]))
		}

		for j := 1; j < app.after+1; j++ {
			if v+j <= len(app.input)-1 {
				if app.printLineNum {
					res.WriteString(fmt.Sprintf("%v-%v\n", v+j+1, app.input[v+j]))
				} else {
					res.WriteString(fmt.Sprintf("%v\n", app.input[v+j]))
				}
			}
		}

		res.WriteString("--\n")
	}

	out := strings.TrimSuffix(res.String(), "--\n")
	out = strings.TrimSpace(out)
	return out
}

func (app *App) findMatches(r *regexp.Regexp) []int {
	matches := make([]int, 0, len(app.input))

	for i, v := range app.input {
		if r.MatchString(v) && !app.invert ||
			!r.MatchString(v) && app.invert {
			matches = append(matches, i)
		}
	}

	return matches
}
