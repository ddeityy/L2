package sorter

import (
	"slices"
	"sort"
	"strings"
)

func (app *App) sort(data []string) []string {
	if app.Reverse {
		slices.Sort(data)
		slices.Reverse(data)
	} else {
		slices.Sort(data)
	}

	if app.deleteDupes {
		data = delDupes(data)
	}

	return data
}

func (app *App) sortColumns(data []string) []string {
	table := Table{
		data:      make([][]string, 0, len(data)),
		column:    app.column - 1,
		isNumeric: app.Numeric,
	}

	for _, v := range data {
		table.data = append(table.data, strings.Fields(v))
	}

	if app.Reverse {
		sort.Sort(sort.Reverse(table))
	} else {
		sort.Sort(table)
	}

	for i, v := range table.data {
		data[i] = strings.Join(v, " ")
	}

	if app.deleteDupes {
		data = delDupes(data)
	}

	return data
}
