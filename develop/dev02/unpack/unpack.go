package unpack

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Runs the CLI app and returns exit status.
func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	input string
}

func (app *appEnv) fromArgs(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "no string provided")
		return fmt.Errorf("no string provided")
	}

	app.input = args[0]
	return nil
}

func (app *appEnv) run() error {
	res, err := (unpack(app.input))
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func unpack(input string) (string, error) {
	runes := []rune(input)
	var b strings.Builder

	for i, v := range runes {
		if unicode.IsDigit(runes[i]) {
			if i == 0 {
				return "", fmt.Errorf("incorrect string")
			}

			letter := runes[i-1]
			if unicode.IsDigit(letter) {
				return "", fmt.Errorf("incorrect string")
			}

			// multi digit unpack a45

			for j := 0; j < int(v-'1'); j++ {
				b.WriteRune(letter)
			}

			continue
		}
		_, err := b.WriteRune(runes[i])
		if err != nil {
			return "", err
		}
	}

	return b.String(), nil
}
