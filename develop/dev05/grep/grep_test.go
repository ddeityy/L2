package grep

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
			},
			want: "Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
		},
		{
			desc: "No matches",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "Primis",
			},
			want: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := regexp.Compile(tC.app.pattern)
			if err != nil {
				t.Error(err)
			}

			got := tC.app.filter(r)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepFixed(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				fixed:   true,
			},
			want: "Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := regexp.Compile(tC.app.pattern)

			if err != nil {
				t.Error(err)
			}

			got := tC.app.filter(r)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepIgnoreCase(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern:    "Primis",
				ignoreCase: true,
			},
			want: "Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.app.pattern = "(?i)" + tC.app.pattern
			r, err := regexp.Compile(tC.app.pattern)

			if err != nil {
				t.Error(err)
			}

			got := tC.app.filter(r)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepInverse(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want []string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				invert:  true,
			},
			want: strings.Fields(`Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.
			--
			Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim.`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := regexp.Compile(tC.app.pattern)
			if err != nil {
				t.Error(err)
			}

			got := strings.Fields(tC.app.filter(r))
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepCount(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				count:   true,
			},
			want: "Matches found: 1",
		},
		{
			desc: "0 matches",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primiz",
				count:   true,
			},
			want: "No matches found",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := regexp.Compile(tC.app.pattern)
			if err != nil {
				t.Error(err)
			}

			got := tC.app.filter(r)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepPrintLineNumbers(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want []string
	}{
		{
			desc: "Normal",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern:      "primis",
				printLineNum: true,
			},
			want: strings.Fields(`2:Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.`),
		},
		{
			desc: "Context",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern:      "primis",
				context:      1,
				printLineNum: true,
			},
			want: strings.Fields(`1-Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.
			2:Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.
			3-Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim.`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.app.after == 0 {
				tC.app.after = tC.app.context
			}
			if tC.app.before == 0 {
				tC.app.before = tC.app.context
			}
			r, err := regexp.Compile(tC.app.pattern)
			if err != nil {
				t.Error(err)
			}

			got := strings.Fields(tC.app.filter(r))
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGrepContext(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		want []string
	}{
		{
			desc: "Before",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				before:  1,
			},
			want: strings.Fields(`Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.
			Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.`),
		},
		{
			desc: "After",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				after:   1,
			},
			want: strings.Fields(`Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.
			Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim.`),
		},
		{
			desc: "Context",
			app: App{
				input: []string{
					"Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.",
					"Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim."},
				pattern: "primis",
				context: 1,
			},
			want: strings.Fields(`Phasellus rutrum eros quis nisi sodales, nec convallis nunc varius. Mauris vehicula ultricies orci, in lacinia ipsum.
			Phasellus in arcu nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus.
			Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam elementum est ac enim pharetra dignissim.`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.app.after == 0 {
				tC.app.after = tC.app.context
			}
			if tC.app.before == 0 {
				tC.app.before = tC.app.context
			}
			r, err := regexp.Compile(tC.app.pattern)
			if err != nil {
				t.Error(err)
			}

			got := strings.Fields(tC.app.filter(r))
			assert.Equal(t, tC.want, got)
		})
	}
}
