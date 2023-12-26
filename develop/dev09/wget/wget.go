package wget

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/schollz/progressbar/v3"
)

type App struct {
	url        *url.URL
	outputFile string
	depth      int
	recursive  bool
}

func CLI(args []string) int {
	var app App

	err := app.fromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Arguement error: %v\n", err)
		return 2
	}

	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}

	return 0
}

func (app *App) fromArgs(args []string) error {
	fl := flag.NewFlagSet("wget", flag.ContinueOnError)
	fl.StringVar(&app.outputFile, "O", "", "Path to output file")
	fl.IntVar(&app.depth, "l", 0, "Recursion depth, maximum is 5")
	fl.BoolVar(&app.recursive, "r", false, "Specify recursive download")

	if err := fl.Parse(args); err != nil {
		return err
	}

	u, err := url.Parse(fl.Arg(0))
	if err != nil {
		return err
	}

	app.url = u

	if app.depth >= 5 {
		return fmt.Errorf("exceeded maximum depth: 5")
	}

	app.depth++

	if app.outputFile == "" {
		app.outputFile = path.Base(app.url.Path)
	}

	return nil
}

func (app *App) run() error {
	if app.recursive {
		queue := []string{app.url.String()}

		if err := os.Mkdir(app.url.Host, os.ModePerm); err != nil {
			return err
		}

		sm := NewSite(app.url.Host, app.url.Host)
		err := sm.DownloadSite(queue, app.depth)
		if err != nil {
			return err
		}

		return nil
	}

	return downloadFile(app.url.String(), app.outputFile)
}

func downloadFile(url string, filePath string) error {
	err := ensureDir(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Downloading %v", filePath),
	)

	size, err := io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s %d bytes\n", filePath, size)

	return nil
}
