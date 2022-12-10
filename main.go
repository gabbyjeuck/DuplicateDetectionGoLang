package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var codeExistsMap = make(map[string]string)

// Get filenames
func getfileNames() []string {
	var files []string
	dataDir := "data/"
	// Load the files within the data directory that end in csv extension
	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

// Load each of the data files
func loadFiles(filename []string) error {
	wg := new(sync.WaitGroup)

	for _, f := range filename {
		wg.Add(1)
		csvFile, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer csvFile.Close()

		csvReader := csv.NewReader(csvFile)
		csvRawData, err := csvReader.ReadAll()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go loadData(csvRawData, f, wg)

	}
	wg.Wait()

	return nil
}

// Load data for each file and validate whether duplicate id exists
func loadData(data [][]string, currentFileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var code string

	for i, c := range data {
		code = c[1]

		if i > 0 {
			// See if existing key for code exists and display original/current filenames where detected.
			if val, exist := codeExistsMap[code]; exist {
				codeExistsMap[code] = val
				fmt.Printf("\nDuplicate exists for ID: %v\nOriginal Filename: %s\nCurrent File: %v \n", code, val, currentFileName)
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			} else {
				codeExistsMap[code] = currentFileName
				continue
			}
		}
	}

}

// Understanding Golang's parallelism features / common golang library functions/interaces.
func main() {
	filenames := getfileNames()
	loadFiles(filenames)
}
