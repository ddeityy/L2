package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "Normal",
			input: "a2df1j4k2l3lnm1mn",
			want:  "aadfjjjjkkllllnmmn",
		},
		{
			desc:  "No numbers",
			input: "abc",
			want:  "abc",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := unpack(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tC.want {
				t.Errorf("got: %s, want: %s", got, tC.want)
			}
		})
	}
}

func TestUnpackOnlyNumbers(t *testing.T) {
	testCases := []struct {
		desc          string
		input         string
		want          string
		expectedError string
	}{
		{
			desc:          "Only numbers",
			input:         "45",
			want:          "",
			expectedError: "invalid string",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := unpack(tC.input)
			assert.EqualError(t, err, tC.expectedError)
			if got != tC.want {
				t.Errorf("got: %s, want: %s", got, tC.want)
			}
		})
	}
}
func TestUnpackEmpty(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "Empty string",
			input: "",
			want:  "",
		},
		{
			desc:  "Single space",
			input: " ",
			want:  " ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := unpack(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tC.want {
				t.Errorf("got: %s, want: %s", got, tC.want)
			}
		})
	}
}

func TestUnpackMultiDigit(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "multiple digit number",
			input: "a12df5cvs13",
			want:  "aaaaaaaaaaaadfffffcvsssssssssssss",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := unpack(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tC.want {
				t.Errorf("got: %s, want: %s", got, tC.want)
			}
		})
	}
}
