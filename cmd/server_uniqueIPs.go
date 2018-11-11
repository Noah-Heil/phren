package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"unicode"

	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

// Page
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// scanResult stores the results of processing a single input file.
type scanResult struct {
	File string   // path to the input file.
	IPs  []net.IP // list of IPs parsed from the file.
	Err  error    // set if an I/O error occurs or the file is empty.
}

// Error satisfies the error interface.
func (r scanResult) Error() string {
	if r.Err == nil {
		return ""
	}
	return fmt.Sprintf("%v: %v", r.File, r.Err)
}

// split is used to divide file content into “words” that might be valid IP
// addresses.
func split(r rune) bool {
	if unicode.IsSpace(r) || unicode.IsPunct(r) && r != '.' && r != ':' {
		return true
	}
	return false
}

// scan reads a file, splits its content in “words,” and tests each word to see
// if it is a valid IPv4 or IPv6 address. If reading the file causes an I/O
// error, or if the file is empty, *scanResult will have a non-nil Err field.
func scan(fp *os.File) *scanResult {
	var (
		res = &scanResult{File: fp.Name()}
		b   []byte
	)
	if b, res.Err = ioutil.ReadAll(fp); res.Err != nil {
		return res
	}
	if len(b) == 0 {
		res.Err = fmt.Errorf("empty file")
		return res
	}
	for _, word := range bytes.FieldsFunc(b, split) {
		if ip := net.ParseIP(string(word)); ip != nil {
			res.IPs = append(res.IPs, ip)
		}
	}
	return res
}

// printError: Just adds some better error formating
func printError(errMsg interface{}) {
	var (
		stderr = colorable.NewColorableStderr()
		red    = color.New(color.FgRed).SprintfFunc()
	)
	fmt.Fprintf(stderr, red("\nerror: %v\n", errMsg))
}

func die(errMsg interface{}) {
	printError(errMsg)
	os.Exit(1)
}

// Extract ips
func findUniqueIPs(targetFiles []string) map[string]int {
	wordFreq := map[string]int{}
	// If any of the input files cannot be read, quit with an error.
	var files []*os.File
	for _, fn := range targetFiles {
		fp, err := os.Open(fn)
		if err != nil {
			die(err)
		}
		files = append(files, fp)
	}

	var (
		results = make(chan *scanResult, len(files))
		wg      sync.WaitGroup
	)

	// Create a goroutine to scan and parse each input file, collecting the
	// results in a channel.
	for _, fp := range files {
		wg.Add(1)
		// TODO: Best practice would be to cap the number of
		// goroutines this can launch, but this is just a simple little
		// tool, and anyone calling it with hundreds of input files
		// should prooooobably already be looking to fork/rewrite this
		// if they want it to be performant.
		go func(fp *os.File) {
			results <- scan(fp)
			fp.Close()
			wg.Done()
		}(fp)
	}

	wg.Wait()
	close(results)

	var failed []*scanResult

	for r := range results {
		// Show successfully extracted IPs first; display errors later.
		if r.Err != nil {
			failed = append(failed, r)
			continue
		}

		for _, ip := range r.IPs {
			wordFreq[ip.String()]++
		}
		fmt.Println()
	}
	if len(failed) > 0 {
		fmt.Println("# errors:")
		for _, r := range failed {
			printError(r)
		}
	}

	return wordFreq

}
