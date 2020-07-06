package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
func main() {

	m := make(map[string]int)

	csvFile, err := os.Open("test.csv")
	defer closeFile(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=country", line[0]))
		defer response.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)

			for _, key := range string(data) {
				m[string(key)] = m[string(key)] + 1
			}
			fmt.Println(string(data))
		}
		fmt.Println(m)
	}

}
