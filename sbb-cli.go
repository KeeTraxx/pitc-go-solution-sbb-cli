package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// Declare command line parameters
	from := flag.String("from", "", "Starting location (e.g. Bern)")
	to := flag.String("to", "", "Destination (e.g. Thun)")
	time := flag.String("time", "", "(optional) Departure time (e.g. hh:mm)")
	date := flag.String("date", "", "(optional) Departure date (e.g. YYYY-MM-DD)")
	help := flag.Bool("help", false, "Display this help screen")

	// Parse command line parameters
	flag.Parse()

	// Check required paramters
	if *help || *from == "" || *to == "" {
		// Print usage if somethings is amiss
		flag.Usage()
		return
	}

	// docs:
	// http://transport.opendata.ch/docs.html#connections
	url := fmt.Sprintf("http://transport.opendata.ch/v1/connections?from=%s&to=%s", *from, *to)

	if *time != "" {
		url = url + fmt.Sprintf("&time=%s", *time)
	}

	if *date != "" {
		url = url + fmt.Sprintf("&date=%s", *date)
	}

	// Make a get request
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// Read the body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	/*var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, body, "", "\t")
	fmt.Println(string(prettyJSON.Bytes()))*/

	// declare connections variable
	var connections ConnectionsResponse

	// Convert json byte array to a go struct
	json.Unmarshal(body, &connections)

	// Loop through all connections and print them
	for _, connection := range connections.Connections {
		fmt.Printf("%-20s %-5s Platform %-5s  =>  %-20s %-5s Platform %-5s\n",
			connection.From.Station.Name,
			connection.From.DepartureTime.Format("15:04"),
			printStrPtr(connection.From.Platform, "n/a"),
			connection.To.Station.Name,
			connection.To.ArrivalTime.Format("15:04"),
			printStrPtr(connection.To.Platform, "n/a"),
		)
	}
}

func printStrPtr(s *string, nilvalue string) string {
	if s == nil {
		return nilvalue
	}
	return *s
}
