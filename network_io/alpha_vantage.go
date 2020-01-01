package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	baseUrl := "https://www.alphavantage.co/query?"
	function := "function=TIME_SERIES_INTRADAY&"
	symbol := "symbol=MSFT&"
	interval := "interval=5min&"
	apiKey := "apikey=" + os.Getenv("ALPHA_API_KEY")

	resp, err := http.Get(baseUrl + function + symbol + interval + apiKey)
	if err != nil {
		log.Println("Initial call")
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println("Parsing JSON")
	//	log.Fatalln(err)
	//}

	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)

	log.Println(body["Meta Data"])
}
