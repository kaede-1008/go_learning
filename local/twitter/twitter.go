package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	// "encoding/json"
	"net/http"
	"html/template"
	"time"
	"regexp"
	"unicode/utf8"
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
}

func getParticipationId(text string) string {
	r := regexp.MustCompile(`^([0-9A-Z]{8})`)
	s := r.FindString(text)

	if s == "" {
		return ""
	}
	return s[0:utf8.RuneCountInString(s)]
}

func Show(writer http.ResponseWriter, request *http.Request) {
	var m []*Tweet
	
	data := struct {
		Title string
	}{
		Title: "Search Key Word",
	}
	// var msg Msg = Msg{}
	// setintervalのような定期実行
	ticker := time.NewTicker(time.Millisecond * 60)
	// slice := []string {"エウロペ", "ワムデュス"}
	// exist := 0

	defer ticker.Stop()
    count := 0
    
	if request.Method == "GET" {
		t, _ := template.ParseFiles("twitter.gtpl")
		t.Execute(writer, data)
	} else {
		request.ParseForm()
		// enemy := request.Form["enemy"]
		// for _, i := range enemy {
		// 	for _, s := range slice {
		// 		if i == s {
		// 			exist++
		// 		}
		// 	}
		// }
		// if len(enemy) != exist {
		// 	msg.Message = msg.Message + "選択肢の中から選んでください"
		// }
		// if msg.Message != "" {
		// 	http.Redirect(writer, request, "/search", 301)
		// }
		word := request.Form["enemy"][0]
		for {
			select {
				case <-ticker.C:
					count++
					
					m = search(word)

					for _, data := range m {
						pid := getParticipationId(data.Text)
						if pid != "" {
							fmt.Fprintln(writer, data.Id)
							fmt.Fprintln(writer, data.Text)
							fmt.Fprintln(writer, pid)
						}
					}
			}
		}	
	}
}

type Tweet struct {
	Text string `json:"text"`
	Id string `json:"id"`
}

type Tweets *[]Tweet

// type Msg struct {
// 	Message string
// }