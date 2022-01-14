package main

import (
	"flag"
	"fmt"
	"os"
	"twitter/creds"
	"twitter/http"
	"twitter/http/client"
	"twitter/http/handler"
	"twitter/marshal"
	"unicode"

	"github.com/olekukonko/tablewriter"
)

func main() {
	creds.Init("./creds/auth.json")

	kwrd := flag.String("query", "TDD", "Search tweets")
	flag.Parse()

	netClient := client.Default()

	hdl := handler.TweetHandler{handler.TweetURLReq{Keyword: *kwrd}, netClient}

	res, err := http.ReqTweets(hdl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("\033[32m")

	renderTweets(res)

	fmt.Print("\033[0m")
}

func renderTweets(tw *marshal.Tweets) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Content", "ID", "Author ID", "Tweeted on"})
	i := 0
	for _, v := range tw.Tweets {
		var cnt string
		for _, v := range v.Content {
			if unicode.Is(unicode.Han, v) {
				//	i += 2
			} else {
				i++
			}

			if i == 40 {
				i = 0
				cnt += "\n"
			}
			cnt += string(v)
		}
		table.Append([]string{cnt, v.Id, v.AuthorId, v.TweetedOn})
	}

	table.SetRowLine(true)
	table.SetBorder(true)
	table.SetCenterSeparator("|")
	table.Render()
}
