package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type dataFileNames struct {
	name string
}

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

// Open the data files
func loadFiles(filename []string) error {
	var code string
	codeExistsMap := make(map[string]bool)
	for _, f := range filename {
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

		for i, c := range csvRawData {
			code = c[1]

			if i > 0 {
				if val, exist := codeExistsMap[code]; exist {
					codeExistsMap[code] = val
					fmt.Printf("Duplicate exists on row %v \n", code)

				} else {
					codeExistsMap[code] = val
					continue
				}
			}
		}
	}
	return nil
}

// Understanding Golang's parallelism features / common golang library functions/interaces.
func main() {
	filenames := getfileNames()
	loadFiles(filenames)
}
