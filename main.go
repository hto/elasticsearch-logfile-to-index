package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/olivere/elastic"
)

var (
	rootPath    = flag.String("rootPath", "/usr/local/temp/elasticIndex", "Root Logs Folder")
	indexName   = flag.String("indexName", "indexName", "ElasticSearch indexName")
	indexType   = flag.String("indexType", "indexType", "ElasticSearch indexType")
	elasticHost = flag.String("elasticHost", "http://127.0.0.1:9200", "Elastic Host")
)

func main() {

	// Run Params
	flag.Parse()

	elasticClient, errElastic := elastic.NewClient(elastic.SetURL(*elasticHost))
	if errElastic != nil {
		panic(errElastic)
	}

	var files []string

	err := filepath.Walk(*rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	fileCount := len(files)
	color.Cyan("File Count : " + strconv.Itoa(fileCount))

	for _, filePath := range files {

		fileCount--

		color.Cyan("File Count : " + strconv.Itoa(fileCount))

		color.Yellow(" - - - - - - " + filePath)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}

		contentType, err := GetFileContentType(file)
		if err != nil {
			continue
		}

		if contentType == "text/plain; charset=utf-8" {

			loopCount := 1
			bulkRequest := elasticClient.Bulk()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {

				bulkRequest = bulkRequest.Add(elastic.NewBulkIndexRequest().Index(*indexName).Type(*indexType).Doc(scanner.Text()))

				// index for every 20000 lines
				if loopCount%20000 == 0 {
					_, err := bulkRequest.Do(context.Background())
					if err != nil {
						fmt.Println(err)
					}
				}

				loopCount++
			}

			_, err = bulkRequest.Do(context.Background())
			if err != nil {
				fmt.Println(err)
			}

		}
		file.Close()
	}
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
