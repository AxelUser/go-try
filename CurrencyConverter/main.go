package main

import(
	"flag"
	"fmt"
)

func main()  {
	var currencyPrt = flag.String("currency", "RUB", "Currency code")
	var valuePtr = flag.Float64("value", 0, "Amount of money to convert")

	fmt.Println(*currencyPrt, *valuePtr)
}