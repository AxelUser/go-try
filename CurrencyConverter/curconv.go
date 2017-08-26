package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Rates for currency
type Rates struct {
	Base       string `json:"base"`
	Date       string `json:"date"`
	Error      string `json:"error"`
	Currencies struct {
		AUD float64 `json:"AUD"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		EUR float64 `json:"EUR"`
		NZD float64 `json:"NZD"`
		RUB float64 `json:"RUB"`
		JPY float64 `json:"JPY"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}

const CurrencyApiBase string = "http://api.fixer.io/latest?base="

func loadRates(currencyCode string) (r *Rates, err error) {
	var rates *Rates

	httpResponse, httpError := http.Get(CurrencyApiBase + currencyCode)
	if httpError != nil {
		log.Println(httpError)
		return nil, errors.New("Could not load rates")
	}
	defer httpResponse.Body.Close()

	body, readError := ioutil.ReadAll(httpResponse.Body)
	if readError != nil {
		log.Println(readError)
		return nil, errors.New("Could not load rates")
	}

	jsonError := json.Unmarshal(body, &rates)
	if jsonError != nil {
		log.Println(jsonError)
		return nil, errors.New("Could not load rates")
	}

	if rates.Error != "" {
		return nil, errors.New("Could not load rates: " + rates.Error)
	}

	return rates, nil
}

func printRates(rates Rates, value float64) {
	fmt.Println("Base currency:\t", rates.Base)
	fmt.Println("Value of:\t", value)

	fmt.Println("\n===== Convert =====\n")

	if rates.Base != "USD" {
		fmt.Println("US Dollar:\t", rates.Currencies.USD*value)
	}
	if rates.Base != "AUD" {
		fmt.Println("Australian Dollar:\t", rates.Currencies.AUD*value)
	}
	if rates.Base != "CAD" {
		fmt.Println("Canadian Dollar:\t", rates.Currencies.CAD*value)
	}
	if rates.Base != "CHF" {
		fmt.Println("Swiss Franc:\t", rates.Currencies.CHF*value)
	}
	if rates.Base != "EUR" {
		fmt.Println("Euro:\t", rates.Currencies.EUR*value)
	}
	if rates.Base != "RUB" {
		fmt.Println("Russian Ruble:\t", rates.Currencies.RUB*value)
	}
	if rates.Base != "JPY" {
		fmt.Println("Japanese Yen:\t", rates.Currencies.JPY*value)
	}
	if rates.Base != "NZD" {
		fmt.Println("New Zealand Dollar:\t", rates.Currencies.NZD*value)
	}
}

func main() {
	var (
		currencyBase string
		value        float64
	)

	flag.StringVar(&currencyBase, "currency", "RUB", "Currency code")
	flag.Float64Var(&value, "value", 1, "Amount of money to convert")

	flag.Parse()

	rates, err := loadRates(currencyBase)
	if err != nil {
		log.Fatal(err)
	}

	printRates(*rates, value)
}
