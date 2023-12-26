package wget

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Site struct {
	root        string
	visitedURLs map[string]struct{}
	directory   string
}

func NewSite(root string, directory string) Site {
	v := map[string]struct{}{
		root: {},
	}

	s := Site{
		root:        root,
		visitedURLs: v,
		directory:   directory,
	}

	return s
}

// Recursively visits links in queue and downloads them.
func (s *Site) DownloadSite(queue []string, depth int) error {
	if depth == 0 {
		return nil
	}

	discoveredLinks := make([]string, 0, 1024)

	for _, v := range queue {
		resp, err := http.Get(v)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		mediatype, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			fmt.Printf("can't parse link type: %s", err.Error())
		}

		ext, err := mime.ExtensionsByType(mediatype)
		if err != nil || len(ext) == 0 {
			fmt.Printf("can't parse link type: %s", err.Error())
			ext = append(ext, "")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var fileName string
		if resp.Request.URL.Path == "/" {
			fileName = path.Join(s.directory, "index"+ext[1])
		} else {
			fileName = path.Join(s.directory, resp.Request.URL.Path+ext[1])
		}

		err = ensureDir(fileName)
		if err != nil {
			return err
		}

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		r := bytes.NewReader(body)

		size, err := io.Copy(file, r)
		if err != nil {
			return err
		}

		fmt.Printf("\nDownloaded %s %d bytes\n", fileName, size)

		if _, err = r.Seek(0, 0); err != nil {
			return err
		}

		l, err := s.parseLinks(r)
		if err != nil {
			return err
		}
		discoveredLinks = append(discoveredLinks, l...)
	}
	if len(discoveredLinks) > 0 {
		depth--
		return s.DownloadSite(discoveredLinks, depth)
	}

	return nil
}

func ensureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return merr
		}
	}

	return nil
}

// Returns a slice of links
func (s *Site) parseLinks(r io.Reader) ([]string, error) {
	res, err := ParseHTML(r)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)

	for _, v := range res {
		href := v.Href.Host + v.Href.Path
		visited := true
		switch {
		case strings.HasPrefix(href, s.root):
			visited = s.isVisited(href)
		case strings.HasPrefix(href, "/"):
			href = s.root + href
			visited = s.isVisited(href)
		}

		if !visited {
			links = append(links, v.Href.Scheme+"://"+href)
		}
	}

	return links, nil
}

// isVisited checks if url visited
func (s *Site) isVisited(href string) (visited bool) {
	if _, visited = s.visitedURLs[href]; !visited {
		s.visitedURLs[href] = struct{}{}
	}

	return visited
}
