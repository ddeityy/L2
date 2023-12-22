package grep

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

type App struct {
	after        int
	before       int
	context      int
	count        bool
	ignoreCase   bool
	invert       bool
	fixed        bool
	printLineNum bool
	reader       io.ReadCloser
	input        []string
	pattern      string
}

func (app *App) FromArgs(args []string) error {
	fl := flag.NewFlagSet("sortfile", flag.ContinueOnError)
	fl.IntVar(&app.after, "A", 0, "print n lines after")
	fl.IntVar(&app.before, "B", 0, "print n lines before")
	fl.IntVar(&app.context, "C", 0, "print n lines before and after")
	fl.BoolVar(&app.count, "c", false, "print a count of matching lines")
	fl.BoolVar(&app.ignoreCase, "i", false, "ignore case")
	fl.BoolVar(&app.invert, "v", false, "exclude instead of matching")
	fl.BoolVar(&app.fixed, "F", false, "match a fixed string instead of a pattern")
	fl.BoolVar(&app.printLineNum, "n", false, "print the line number of a match")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	if len(fl.Args()) == 0 {
		fl.Usage()
		os.Exit(1)
	}

	if app.after == 0 {
		app.after = app.context
	}

	if app.before == 0 {
		app.before = app.context
	}

	app.pattern = fl.Arg(0)

	if app.ignoreCase {
		app.pattern = "(?i)" + app.pattern
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}

	file, err := os.Open(fl.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", fl.Arg(1), err)
		return err
	}
	app.reader = file

	return nil
}

func (app *App) Grep() (string, error) {
	defer app.reader.Close()
	scanner := bufio.NewScanner(app.reader)
	for scanner.Scan() {
		app.input = append(app.input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	r, err := regexp.Compile(app.pattern)
	if err != nil {
		return "", err
	}

	res := app.filter(r)

	return res, nil
}
