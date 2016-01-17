package main

import (
	"fmt"
	"os"

	"github.com/MDrollette/taxjar-go"
)

func main() {
	c := taxjar.NewClient(os.Getenv("API_TOKEN"))
	c.Debug = true

	// Get rates at specific ZIP with an optional city specifier
	// rate, err := c.Rates.Get("12901", taxjar.RateCity("Plattsburgh"))
	// if nil != err {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%+v\n", rate)

	taxes, err := c.Taxes.Calculate("2211 Commerce St.", "Dallas", "TX", "75201", "US", "TX", "75206", "US", 100.00, 10.00)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", taxes)
}
