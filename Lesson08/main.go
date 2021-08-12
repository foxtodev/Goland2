package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	removeDup bool   = false
	dirScan   string = ""
	isDup     bool   = false
	wg        sync.WaitGroup
)

// Result structure
type Result struct {
	file string
	size int64
}

func worker(input chan string, results chan<- *Result) {
	wg.Add(1)
	defer wg.Done()
	for file := range input {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fi, err := f.Stat()
		if err != nil {
			panic(err)
		}
		defer func() {
			// recover from panic caused by writing to a closed channel
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
				fmt.Printf("write: error writing %s on channel: %v\n", fi.Name(), err)
				return
			}
		}()
		results <- &Result{
			file: file,
			size: fi.Size(),
		}
	}
}

func findAllFiles(input chan string) {
	filepath.Walk(dirScan, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else if info.Mode().IsRegular() {
			input <- path
		}
		return nil
	})
	close(input)
}

func main() {
	// FLAGS
	flag.StringVar(&dirScan, "dir", "temp", "Directory")
	flag.BoolVar(&removeDup, "remove", false, "Delete all duplicates")
	var help = flag.Bool("h", false, "Display this message")
	var helpp = flag.Bool("help", false, "Display this message")
	flag.Parse()

	if *help || *helpp {
		fmt.Println("\nduplicates is a command line tool to find duplicate files in a folder")
		fmt.Println("\nusage: duplicates [options...] path")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("\nfind duplicate files in a folder", dirScan)

	input := make(chan string)
	results := make(chan *Result, 10)

	workers := 1000 //runtime.NumCPU()

	// wg := sync.WaitGroup{}
	// wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(input, results) //, &wg)
	}

	go findAllFiles(input)
	go func() {
		wg.Wait()
		close(results)
	}()

	counter := make(map[Result][]string)
	for result := range results {
		fileNameSize := Result{filepath.Base(result.file), result.size}
		counter[fileNameSize] = append(counter[fileNameSize], result.file)
		if len(counter[fileNameSize]) > 1 {
			isDup = true
		}
	}

	if !isDup {
		fmt.Println("no duplicate files were found in the folder")
		os.Exit(0)
	}

	for files, duplicates := range counter {
		if len(duplicates) > 1 {
			fmt.Printf("\nFound %d duplicates for %s (%d bytes): \n", len(duplicates), files.file, files.size)
			for i, f := range duplicates {
				fmt.Print(" ", f)
				if removeDup && i != 0 {
					err := os.Remove(f)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(" Removed")
				} else {
					fmt.Println("")
				}
			}
		}
	}

	if !removeDup && isDup {
		var b []byte = make([]byte, 1)
		fmt.Print("\nwould you like to delete duplicates? (Y/...) ")
		os.Stdin.Read(b)
		if string(b) == "Y" {
			fmt.Println("Removing...")
			for _, duplicates := range counter {
				if len(duplicates) > 1 {
					for i, f := range duplicates {
						if removeDup && i != 0 {
							err := os.Remove(f)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}
			}
		}
		os.Exit(0)
	}

	fmt.Println("\nall duplicate files are removed")

}
