package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"os"
	"log"
	"io"
	"encoding/json"
	"encoding/csv"
)

type Sheet struct{
	Sitename string `json:"sitename"`
	Location string `json:"location"`
	Address *Address `json:address`

}

type Address struct {
	Number string `json:"number"`
}



func main() {

	csvFile, _ := os.Open("test.csv")
	reader:= csv.NewReader(csvFile)
	var sheet [] Sheet
	for {
	 line, error:= reader.Read()
		if error==io.EOF {
			break
		} else if error!= nil{
			log.Fatal(error)
		}
		response, err:= http.Get("http://ip-api.com/json/www.google.com?fields=country")
		/*I  want to able to pass line[0] which is the website in my test csv in the url*
		Something like http.Get("http://ip-api.com/json/{line[0}?fields=country") only returns empty JSON {}*/
		if err!= nil{
		fmt.Println("HTTP request failed with error")
		} else{	
			data, _:= ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	
  
		sheet = append(sheet, Sheet{
			Sitename: line[0],
			Location: line[1],
			Address: &Address{
				Number:line[2],
			},

		})
	}
	sheetjson, _:= json.MarshalIndent(sheet, "", "  ")
	fmt.Println(string(sheetjson))
}
