package sorter

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type App struct {
	Numeric     bool
	Reverse     bool
	deleteDupes bool
	column      int
	reader      io.ReadCloser
}

func (app *App) FromArgs(args []string) error {
	fl := flag.NewFlagSet("sortfile", flag.ContinueOnError)
	fl.IntVar(&app.column, "k", 1, "sort via a key")
	fl.BoolVar(&app.Numeric, "n", false, "compare according to string numerical value")
	fl.BoolVar(&app.Reverse, "r", false, "reverse the result of comparisons")
	fl.BoolVar(&app.deleteDupes, "u", false, "delete duplicate strings")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}

	file, err := os.Open(fl.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", fl.Arg(0), err)
		return err
	}
	app.reader = file

	return nil
}

func (app *App) Run() error {
	defer app.reader.Close()
	data := make([]string, 0)

	scanner := bufio.NewScanner(app.reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if app.column == 1 && !app.Numeric {
		data = app.sort(data)
		for _, v := range data {
			fmt.Fprintf(os.Stdout, "%s\n", v)
		}
		return nil
	}

	data = app.sortColumns(data)
	for _, v := range data {
		fmt.Fprintf(os.Stdout, "%s\n", v)
	}

	return nil
}
