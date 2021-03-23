package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"encoding/json"
	// "testing"
	// "github.com/stretchr/testify/assert"
)



func TestConnectTwitterApi() *anaconda.TwitterApi {
	// 環境変数からkeyの取得
	accessToken := os.Getenv("accessToken")
	accessTokenSecret := os.Getenv("accessTokenSecret")
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	fmt.Println(accessToken)

	// 認証
	return anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)

}

func TestSearch(word string) {
	api := TestConnectTwitterApi()
	fmt.Println(api)
	// 検索
	searchResult, _ := api.GetSearch(`"` + word + `"`, nil)
	fmt.Println(searchResult)
	tweets := make([]*Tweet, 0)

	for _, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.User = data.User.Name
		tweet.Text = data.FullText
		tweets = append(tweets, tweet)
	}
	s, err := json.Marshal(tweets)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}



// func Twitter() {
// 	// 環境変数からkeyの取得
// 	accessToken := os.Getenv("accessToken")
// 	accessTokenSecret := os.Getenv("accessTokenSecret")
// 	consumerKey := os.Getenv("consumerKey")
// 	consumerSecret := os.Getenv("consumerSecret")

// 	// 認証
// 	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)

// 	// 検索
// 	searchResult, _ := api.GetSearch(`"グラブル"`, nil)
// 	for _, tweet := range searchResult.Statuses {
// 		fmt.Println(tweet.Text)
// 	}
// }

type Tweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type Tweets *[]Tweet