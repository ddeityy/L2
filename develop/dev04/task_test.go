package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		desc string
		have []string
		want map[string][]string
	}{
		{
			desc: "english",
			have: []string{
				"angel", "angle", "glean", "alert", "alter", "later", "AngEl", "ALERT", "algol",
			},
			want: map[string][]string{
				"angel": {"angle", "glean"},
				"alert": {"alter", "later"},
			},
		},
		{
			desc: "russian",
			have: []string{
				"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик", "стоЛик", "листОК",
			},
			want: map[string][]string{
				"листок": {"слиток", "столик"},
				"пятак":  {"пятка", "тяпка"},
			},
		},
		{
			desc: "chinese",
			have: []string{
				"航合投職羽張123", "羽12張3航合投職", "2張1合職3羽航投", "21羽合張投3航職", "投張羽合2職31航", "航張職2投13合羽", "航張職2投13合羽", "航張職2投13合羽",
			},
			want: map[string][]string{
				"航合投職羽張123": {"羽12張3航合投職", "2張1合職3羽航投", "21羽合張投3航職", "投張羽合2職31航", "航張職2投13合羽"},
			},
		},
		{
			desc: "no anagrams",
			have: []string{
				"123f", "ghskjas", "sfasdfk",
			},
			want: map[string][]string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findAnagrams(tC.have)
			assert.Equal(t, tC.want, got)
		})
	}
}
