package mbase

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
)

var DATA_DIR = os.Getenv("DATA_DIR")

var INDEX_LIST = OpenIndexes()

func OpenIndexes() map[string]bleve.Index {

	IndexList := make(map[string]bleve.Index)

	if DATA_DIR == "" { // Data directory does not in env variable
		DATA_DIR = "data"
		if _, err := os.Stat("data"); os.IsNotExist(err) { // check if data directory exists anyways
			err := os.Mkdir("data", 0755) // create data directory if does not exist
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created data directory: data")
		}
	}

	// walk the data dir and register index names
	dirEntries, err := ioutil.ReadDir(DATA_DIR)
	if err != nil {
		log.Fatalf("error reading data dir: %v", err)
	}

	for _, dirInfo := range dirEntries {
		indexPath := DATA_DIR + string(os.PathSeparator) + dirInfo.Name()

		// skip single files in data dir since a valid index is a directory that
		// contains multiple files
		if !dirInfo.IsDir() {
			log.Printf("not registering %s, skipping", indexPath)
			continue
		}

		i, err := bleve.Open(indexPath)
		if err != nil {
			log.Printf("error opening index %s: %v", indexPath, err)
		} else {
			log.Printf("registered index: %s", indexPath)
			// set correct name in stats
			i.SetName(dirInfo.Name())
			IndexList[dirInfo.Name()] = i
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return IndexList

}
