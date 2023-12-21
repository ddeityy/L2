package sorter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "normal",
			app:  App{},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
		{
			desc: "Not numeric",
			app:  App{},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sort(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortDupes(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "Delete duplicates",
			app: App{
				deleteDupes: true,
			},
			data: []string{
				"1",
				"1",
				"5",
				"13",
				"23",
				"11",
				"11",
				"21",
				"31",
				"31",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sort(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortReverse(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "Reverse order",
			app: App{
				Reverse: true,
			},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"5",
				"31",
				"23",
				"21",
				"13",
				"11",
				"1",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sort(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumnsOutOfRange(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "Column out of range",
			app: App{
				column: 200,
			},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortColumns(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumnsNumWithLetter(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "Numeric, but column starts with a letter",
			app: App{
				column:  1,
				Numeric: true,
			},
			data: []string{
				"d1",
				"ad5",
				"asbv13",
				"sfg23",
				"fa11",
				"gh21",
				"31",
			},
			want: []string{
				"31",
				"ad5",
				"asbv13",
				"d1",
				"fa11",
				"gh21",
				"sfg23",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortColumns(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumns(t *testing.T) {
	testCases := []struct {
		desc string
		app  App
		data []string
		want []string
	}{
		{
			desc: "Normal",
			app: App{
				column: 2,
			},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"I am my aunt's sister's daughter.",
				"Cats are good pets, for they are clean and are not noisy.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Standing on one's head at job interviews forms a lasting impression.",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortColumns(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
