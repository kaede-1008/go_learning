package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	// "encoding/json"
	"net/http"
	"html/template"
	// "log"
	// simplejson "github.com/bitly/go-simplejson"
)

func connectTwitterApi() *anaconda.TwitterApi {
	// 環境変数からkeyの取得
	AccessToken := os.Getenv("accessToken")
	AccessTokenSecret := os.Getenv("accessTokenSecret")
	ConsumerKey := os.Getenv("consumerKey")
	ConsumerSecret := os.Getenv("consumerSecret")


	// 認証
	return anaconda.NewTwitterApiWithCredentials(AccessToken, AccessTokenSecret, ConsumerKey, ConsumerSecret)

}

func search(word string) []*Tweet {
	if os.Getenv("accessToken") == "" {
		Set()
	}
	api := connectTwitterApi()

	// 検索
	searchResult, _ := api.GetSearch(`"` + word + `"`, nil)
	tweets := make([]*Tweet, 0)

	for _, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.Text = data.FullText
		tweet.Id = data.IdStr
		tweets = append(tweets, tweet)
	}
	return tweets
	// out, err := json.Marshal(tweets)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(string(out))
	// return out
}

func Show(writer http.ResponseWriter, request *http.Request) {
	var m []*Tweet

	if request.Method == "GET" {
		t, _ := template.ParseFiles("twitter.gtpl")
		t.Execute(writer, nil)
	} else {
		word := request.FormValue("word")

		m = search(word)

		// err := json.Unmarshal(s_result, &m)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(m)
		for _, data := range m {
			fmt.Fprintln(writer, data.Id)
			fmt.Fprintln(writer, data.Text)
		}

		// json, err := simplejson.NewJson(s_result)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		
		// for _, datas := range json.MustArray() {
		// 	// fmt.Fprintf(writer, json.GetIndex(i))
		// 	// fmt.Println(json.GetIndex(i))
		// 	fmt.Println(datas["text"])

		// }
	}
}

type Tweet struct {
	Text string `json:"text"`
	Id string `json:"id"`
}

type Tweets *[]Tweet