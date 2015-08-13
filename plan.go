package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jhkennedy4/go-qpx"
)

var API_URL string = "https://www.googleapis.com/qpxExpress/v1/trips/search?key=" + os.Getenv("QPX_KEY")

const (
	origin = iota
	destination
	start_date
	end_date
)

func main() {
	args := os.Args[1:]

	var args_length int = len(args)
	// usage
	if args_length < 4 {
		fmt.Println("Usage: plan origin destination start_date end_date")
		// [vacations_days] [start_date] [end_date] (format: YYYY-MM-DD[ HH[:MM[:SS]]])")
		return
	}

	query := qpx.NewQuery(args[origin], args[destination], args[start_date], args[end_date])

	fmt.Println(args)

	resp, err := http.Post(API_URL, "application/json", query.ToJSON())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s", body)
	r := qpx.ParseResponse(body)

	fmt.Printf("%s", r.PrettyPrint())

}
