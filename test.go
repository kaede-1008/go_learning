package main

import (
	"fmt"
	"glablueProject/local/twitter"
)

func main() {
	twitter.Set()
	tweets := twitter.Search(`美味しい`)
	fmt.Println(tweets)
}