package filtering

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-void/portal/pkg/filter"
)

type FilterWatcher struct {
	// WatchURLS indicates if URL watching is enabled
	WatchURLS bool

	// WatchFiles indicates if file watching is enabled
	WatchFiles bool

	// WatchInterval defines the watch interval in which URLs should
	// be checked for changes
	WatchInterval int

	// Files stores a slice of file paths
	Files []string

	// Urls stores a slice of URLs
	Urls []string

	Filter filter.Filter
}

func (w *FilterWatcher) AddRulesFromFile(t filter.RuleType, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(file)

	for s.Scan() {
		err := w.Filter.AddRule(t, s.Text())
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *FilterWatcher) AddRulesFromURL(t filter.RuleType, url string) error {
	// TODO (Techassi): Don't use default HTTP client, use custom one so that the user can adjust options
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("filter: request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("filter: failed to read body: %v", err)
	}

	r := bytes.NewReader(body)
	s := bufio.NewScanner(r)

	for s.Scan() {
		err := w.Filter.AddRule(t, s.Text())
		if err != nil {
			return err
		}
	}

	return nil
}
