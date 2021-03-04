package twitter

import (
	// "fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	// "net/http"
)

func connectTwitterApi() *anaconda.TwitterApi {
	// 環境変数からkeyの取得
	accessToken := os.Getenv("accessToken")
	accessTokenSecret := os.Getenv("accessTokenSecret")
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")

	// 認証
	return anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)

}

func Search(word string) []string {
	api := connectTwitterApi()

	// 検索
	searchResult, _ := api.GetSearch(`"` + word + `"`, nil)

	tweets := make([]string, 0)

	for _, data := range searchResult.Statuses {

		tweets = append(tweets, data.Text)
	}

	return tweets
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
