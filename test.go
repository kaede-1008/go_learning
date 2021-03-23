package main

import (
	// "fmt"
	"glablueProject/local/twitter"
)

func main() {
	twitter.Set()
	twitter.TestSearch(`美味しい`)
	// fmt.Println(tweets)
}