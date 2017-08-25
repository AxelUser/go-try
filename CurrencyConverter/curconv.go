package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

func loadRates(currencyCode string) Rates {
	var rates Rates

	httpResponse, httpError := http.Get(CurrencyApiBase + currencyCode)
	if httpError != nil {
		fmt.Println(httpError)
	}
	defer httpResponse.Body.Close()

	body, readError := ioutil.ReadAll(httpResponse.Body)
	if readError != nil {
		fmt.Println(readError)
	}

	jsonError := json.Unmarshal(body, &rates)
	if jsonError != nil {
		fmt.Println(jsonError)
	}

	return rates
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

	rates := loadRates(currencyBase)

	if rates.Error == "" {
		printRates(rates, value)
	} else {
		fmt.Print("Error:\t", rates.Error)
	}
}