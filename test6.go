package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
func main() {

	f := flag.String("f", "", "Path to input file")
	out := flag.String("out", "", "Path to output file")
	n := flag.Int("n", 0, "Number of websites to analyze")
	flag.Parse()
	filee := *f
	num := *n
	outt := *out

	m := make(map[string]string)
	r := make(map[string]string)

	counter := 1

	csvFile, err :=
		os.Open(filee)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(csvFile)
	reader := csv.NewReader(csvFile)
	for i := 1; i <= num; i++ {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=country", line[0]))

		if err != nil {
			fmt.Println(err)
			defer response.Body.Close()

		} else {

			data, _ := ioutil.ReadAll(response.Body)
			err := json.Unmarshal(data, &m)

			for k := range m {
				con := fmt.Sprint("", counter)
				i := m[k] // m[k] is the value "Singapore" and k is "country" only need first value

				if _, ok := r[i]; !ok { //if key doesnt exist initialize counter to 1
					r[i] = con // r[i] is the value "number" and i the key is "Singapore"

				} else {
					newVal, _ := strconv.Atoi(r[i]) // if key exists increase by 1
					newVal++
					r[i] = fmt.Sprint("", newVal)
				}
			}

			if err != nil {
				panic(err)
			}

		}

	}
	file, fileErr := os.Create(outt)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	test, _ := json.MarshalIndent(r, "", "")
	fmt.Fprint(file, string(test))

}
