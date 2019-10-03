package main

import (
	"fmt"
	"log"
	"os"

	"github.com/singmyr/centra-client-shop"
)

func main() {
	fmt.Println("Centra client test")
	client := centra.Init(os.Getenv("CENTRA_URL"), os.Getenv("CENTRA_SECRET"))

	markets, _ := client.GetMarkets()
	for k, v := range markets {
		fmt.Println("Market", k, v.Name)
	}

	market, e := client.GetMarket("27")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(market.Name)

	// fmt.Printf("%# v", pretty.Formatter(market))
}
